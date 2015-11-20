var fs = require("fs")
var path = require('path');
var ejs = require('ejs');
var removeExistingFiles = true;

var metaDir = "../openapi-meta-master/api-metadata/";

var goCodeDir = "../src/"

//response modal
var serviceTemplate = fs.readFileSync(__dirname + '/service.ejs', 'utf8');

var requestModalTemplate = fs.readFileSync(__dirname + '/modal_req.ejs', 'utf8');
var transoformerWithBuilder = fs.readFileSync(__dirname + '/transformer.ejs', 'utf8');

var urlPattern = /\[(.*?)\]/g;

fs.readdir(metaDir, function (err, files) {

    if (err == null) {
        files.forEach(function (curVal, idx) {
            if (curVal[0] != ".") {
                createServiceCode(curVal);
            }
        });
    }
    else {
        err(err);
    }
});


function readJson(jsonFile) {

    var jsonContent = fs.readFileSync(jsonFile, 'utf8');
    jsonContent = jsonContent.replace(/\r?\n|\r|\t/gm, "");

    return JSON.parse(jsonContent);;
}

function createServiceCode(serviceDir) {

    var serviceFullDir = metaDir + serviceDir;

    var serviceSubFolders = fs.readdirSync(serviceFullDir);
    serviceSubFolders = serviceSubFolders.filter(function (item) {
        return item[0] != ".";
    });
    serviceSubFolders.sort()

//Read date folder
    for (var i = 0; i < serviceSubFolders.length; i++) {

        var serviceSrcDir = path.join(goCodeDir, serviceDir.replace("api-metadata", "go-sdk"), serviceSubFolders[i].replace(/-/g, ""))
        var serviceShortName = serviceDir.split("-").reverse()[0]
        
        // var packageName = serviceDir.replace(/-/g, ".") + ".v" + serviceSubFolders[i].replace(/-/g,"");
        var packageName = serviceShortName;
        
        //create service srouce dir
        mkdirSync(serviceSrcDir);

        var serviceMetaDirRoot = path.join(serviceFullDir, serviceSubFolders[i]);
        var verInfo = readJson(path.join(serviceMetaDirRoot, "Version-Info.json"));

        var serviceName = capitalizeFirstLetter(serviceShortName) + "Client";

        createServiceSrc(serviceName, packageName, verInfo, serviceSrcDir, serviceMetaDirRoot)

    }

}

function rebuildRequestParasObject(apiInfo) {
    if (!apiInfo) {
        return null;
    }

    var paras = apiInfo.parameters.parameters;

    var newParaObj = {};
    
    if(apiInfo.isvProtocol){
	   if(apiInfo.isvProtocol.method)
        newParaObj.RequestMethod = apiInfo.isvProtocol.method;
		
		if(apiInfo.isvProtocol.protocol)
		   newParaObj.ReqProtocol =apiInfo.isvProtocol.protocol;
    }
    else{
        newParaObj.RequestMethod = "GET";
		newParaObj.ReqProtocol= "Http";
    }
     
    newParaObj.properties = apiInfo.parameters.parameters;

    if (apiInfo.isvProtocol.pattern) {
        var newPattern = apiInfo.isvProtocol.pattern;
        var propNames = [];

        while ((m = urlPattern.exec(apiInfo.isvProtocol.pattern)) !== null) {
            if (m.index === urlPattern.lastIndex) {
                urlPattern.lastIndex++;
            }

            newPattern = newPattern.replace(m[0], "%s");
            propNames.push(m[1]);
        }

        newParaObj.pattern = newPattern;
        var propCodeLst = "";

        for (var idx in propNames) {
            propCodeLst += ",r." + propNames[idx];
        }
        newParaObj.patternProps = propCodeLst;
    }

    newParaObj.Position = {};

    for (var index = 0; index < paras.length; index++) {
        var element = paras[index];

        if (newParaObj.Position[element.tagPosition] == undefined) {
            newParaObj.Position[element.tagPosition] = [];
        }

        setReqValueFuncName(element);
        newParaObj.Position[element.tagPosition].push(element);
    }

    if (newParaObj.Position["Query"] || newParaObj.Position["Body"]) {
        newParaObj.hasProp = 1
    }

    if (newParaObj.Position["Path"]) {
        newParaObj.hasPath = 1
    }

    return newParaObj;

}

function createServiceSrc(serviceName, packageName, versionInfo, serviceSrcDir, serviceMetaDirRoot) {

    var methodInfos = [];
    versionInfo.apis.apis.forEach(function (curVal, idx) {
        var apiFileName = path.join(serviceMetaDirRoot, "Api", curVal.name + ".json");

        if (fs.existsSync(apiFileName)) {

            var apiInfo = null;

            try {
                apiInfo = readJson(apiFileName)

            }
            catch (e) {
                err("Wrong jason format:"+apiFileName)
            }

            try {
                if (apiInfo != null) {
                    
                    // if(curVal.name == "DescribeCdnDomainLogs"){
                    //     debugger;
                    // }
                    var requestParaLen = apiInfo.parameters.parameters.length;
                    var reseponseParaLen = (apiInfo.resultMapping.members ? apiInfo.resultMapping.members.length : 0) +
                        + (apiInfo.resultMapping.structs ? apiInfo.resultMapping.structs.length : 0)
                        + (apiInfo.resultMapping.arrays ? apiInfo.resultMapping.arrays.length : 0);

                    var methodInfo = {
                        requestParaLen: requestParaLen,
                        reseponseParaLen: reseponseParaLen,
                        detail: curVal
                    };
                    
                    //using for generate service code
                    methodInfos.push(methodInfo);
          
                    //-------------write modal start---------------
                    createServiceRequestResponseModal(packageName, apiInfo, serviceSrcDir, requestParaLen, reseponseParaLen);
                    //-------------write modal end---------------
          
         
                } else {
                    err("Wrong json format file:" + apiFileName);
                }
            } catch (e) {
                err("Unknow error:" + apiFileName);
                throw e;
            }
            
            //console.log("Done:" + apiFileName);

        }
        else {
            err("Missing file:" + apiFileName);
        }


    });
       
    var absServiceDir = serviceSrcDir.replace('..\\src\\','').replace("\\","/");   
    //-------------write servicode start---------------
    var serviceCode = ejs.render(serviceTemplate, {
        packageName: packageName,
        absServicePath:absServiceDir,
        service: serviceName,
        Version: versionInfo.name,
        Product: versionInfo.product,
		ApiStyle:versionInfo.apiStyle,
        methodInfos: methodInfos
    });



    writeFile(path.join(serviceSrcDir, serviceName + ".go"), serviceCode);
    //------------write servicode end--------------  
}

function log(msg) {

}

function err(msg) {
    console.log(msg)
}

function upperFirst(name) { return name[0].toUpperCase() + name.substr(1); }

function createServiceRequestResponseModal(packageName, apiInfo, serviceSrcDir, requestParaLen, reseponseParaLen) {

    //------------------Request modal start----------------------
    var newReqModal = rebuildRequestParasObject(apiInfo);
    replateBasicTypeName(newReqModal.properties);
    var modalName = upperFirst(apiInfo.name) + "Request";
    var reqModalCode = ejs.render(requestModalTemplate, {
        actionName: apiInfo.name,
        packageName: packageName,
        modalName: modalName,
        RequestMethod:newReqModal.RequestMethod,
        modalInfo: newReqModal//apiInfo.parameters.parameters
    });

    writeFile(path.join(serviceSrcDir, modalName + ".go"), reqModalCode);
    //------------------Request modal end----------------------

    var memberInfo = apiInfo.resultMapping;
    memberInfo.name = apiInfo.name + "Response";

    if (!memberInfo.members) {
        memberInfo.members = [];
    }

    memberInfo.members.push({ tagName: "RequestId", "type": "string" });

    var responseDir = path.join(serviceSrcDir,upperFirst(apiInfo.name) + "Response");
    mkdirSync(responseDir);
    writeServiceModalCode("",apiInfo.resultMapping,packageName +"_"+upperFirst(apiInfo.name) + "Response",apiInfo.name + "Response",responseDir);
   // handlObjectMemberinfo(memberInfo, serviceSrcDir, packageName, memberInfo.members);
    //writeModalCode(serviceSrcDir, { packageName: packageName, modalName: memberInfo.name, members: memberInfo.members });

}

function getPath(parentKey,currentTag){
    if(parentKey == "."){
        return  "." + currentTag;     
    }else{
   return parentKey + "." + currentTag;     
    }
   
}

function renameTagName(tagName){
    return tagName.replace(/\./g,"_")
}

function rebuildServiceModal(parentKey,resultMapping,ret){
    var props = [];

   if(resultMapping.arrays && resultMapping.arrays.length > 0){
    ret.HasArray = true;
      for(idx in resultMapping.arrays){
          props.push({
             "tagName":resultMapping.arrays[idx].tagName,
             "tagNameAlias":renameTagName(resultMapping.arrays[idx].tagName),
             "typeName" : "[]" + upperFirst(resultMapping.arrays[idx].itemName),
             "basicType": resultMapping.arrays[idx].itemName,
             "PropType":0,
             "Key":getPath("",resultMapping.arrays[idx].tagName)
          }); 
      }
   } 
   
   if(resultMapping.members && resultMapping.members.length > 0){
    ret.HasMember = true;
        for(idx in resultMapping.members){
             props.push({
             "tagName":resultMapping.members[idx].tagName,
             "tagNameAlias":renameTagName(resultMapping.members[idx].tagName),
             "typeName" : replateBasicTypeNameSingle(resultMapping.members[idx].type),
             "PropType":1,
             "Key":getPath("",resultMapping.members[idx].tagName),
             "GetValFunc":getValFunName(resultMapping.members[idx].type)
          }); 
        }
   }

   if(resultMapping.structs && resultMapping.structs.length > 0){
    ret.HasStruct = true;
       for(idx in resultMapping.structs){
             props.push({
             "tagName":resultMapping.structs[idx].tagName,
             "tagNameAlias":renameTagName(resultMapping.structs[idx].tagName),
             "typeName" : upperFirst(resultMapping.structs[idx].tagName),
             "PropType":2,
             "Key":getPath("",resultMapping.structs[idx].tagName)
          }); 
        }
   }

   if(resultMapping.lists && resultMapping.lists.length > 0){
    ret.HasList = true;
        for(idx in resultMapping.lists){
             props.push({
             "tagName":resultMapping.lists[idx].tagName,
             "tagNameAlias":renameTagName(resultMapping.lists[idx].tagName),
             "typeName" : "[]string",
             "PropType":3,
             "Key":getPath("",resultMapping.lists[idx].tagName)
          }); 
        }
   }
ret.Props = removeDuplicateKey(props);
   return ret;
}

function writeServiceModalCode(parentKey,resultMapping,packageName,modalName,serviceSrcDir) {
    var codeTemplateObj = {HasArray:false,HasList:false,HasMember:false,HasStruct:false};
    
    codeTemplateObj.packageName = packageName;
    codeTemplateObj.modalName = modalName;

    rebuildServiceModal(parentKey,resultMapping,codeTemplateObj);

    var code = ejs.render(transoformerWithBuilder, codeTemplateObj);
    writeFile(path.join(serviceSrcDir,modalName + ".go"), code);

   if(resultMapping.arrays && resultMapping.arrays.length > 0){
       for(idx in resultMapping.arrays){
          writeServiceModalCode(parentKey+"." + resultMapping.arrays[idx].tagName,
                               resultMapping.arrays[idx],
                               packageName,
                               resultMapping.arrays[idx].itemName,
                               serviceSrcDir);
       }
    }

    if(resultMapping.structs && resultMapping.structs.length > 0){
       for(idx in resultMapping.structs){
          writeServiceModalCode(parentKey + "." + resultMapping.structs[idx].tagName,
                               resultMapping.structs[idx],
                               packageName,
                               resultMapping.structs[idx].tagName,
                               serviceSrcDir);
       }
    }
}



function removeDuplicateKey(properties) {
    var objsHash = {};
    var newProperties = []
    for (var index = 0; index < properties.length; index++) {
        var element = properties[index];

        if (!objsHash[element.tagName]) {
            objsHash[element.tagName] = true;
            newProperties.push(element);
        }

    }

    return newProperties;
}

function writeFile(src, code) {
    
    fs.writeFileSync(src, code);
}

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

function mkdirSync(url, mode, cb) {
    
    //clear all existed source files
    if (fs.existsSync(url)) {

        if (removeExistingFiles) {
            var files = fs.readdirSync(url);
            files.forEach(function (file, index) {
                var curPath = url + "/" + file;
                if (!fs.statSync(curPath).isDirectory()) {
                    fs.unlinkSync(curPath);
                }
            });
        }

        return;
    }
    
    var path = require("path"), arr = url.split("\\");

    mode = mode || 0755;
    cb = cb || function () { };
    if (arr[0] === ".") {//处理 ./aaa
        arr.shift();
    }
    if (arr[0] == "..") {//处理 ../ddd/d
        arr.splice(0, 2, arr[0] + "\\" + arr[1])
    }
    function inner(cur) {
        if (!path.existsSync(cur)) {//不存在就创建一个
            fs.mkdirSync(cur, mode)
        }
        if (arr.length) {
            inner(cur + "\\" + arr.shift());
        } else {
            cb();
        }
    }
    arr.length && inner(arr.shift());
}

function getValFunName(propType) {
        switch (propType.toLowerCase()) {
            case "integer":
                return "builder.GetInt32"
                break;
            case "string":
                return "builder.GetString"
                break;
            case "long":
                return "builder.GetInt64"
                break;
            case "boolean":
                return "builder.GetBool"
            case "float":
                return "builder.GetFloat32"
                break;
            default:
                err("Unknow type:" + propType);
                return undefined;
                break;
        }
}

 


function replateBasicTypeNameSingle(memType) {
    switch (memType.toLowerCase()) {
            case "integer":
                return "int32";
                break;
            case "string":
                return "string";
                break;
            case "long":
                return "int64";
                break;
            case "boolean":
                return "bool";
            case "float":
                return "float32";
                break;
            case "list":
                return "[]string";
                break;

            default:
            console.log("Wrong member type: " +memType);
             return undefined;
                break;
        }
}


function setReqValueFuncName(mem) {
  switch (mem.type.toLowerCase()) {
            case "integer":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueInt32";
                break;
            case "string":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueStr";
                break;
            case "long":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueInt64";
                break;
            case "boolean":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueBoolean";
                 break;
            case "float":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueFloat32";
                break;
            case "list":
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueStrArr";
                break;

            default:
                 mem.GetValFunc = "efg";
                break;
        }
}

function replateBasicTypeName(members) {
    members.forEach(function (mem) {
        switch (mem.type.toLowerCase()) {
            case "integer":
                mem.type = "int32";
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueInt32";
                break;
            case "string":
                mem.type = "string";
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueStr";
                break;
            case "long":
                mem.type = "int64";
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueInt64";
                break;
            case "boolean":
                mem.type = "bool";
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueBoolean";
                 break;
            case "float":
                mem.type = "float32";
                mem.GetValFunc = "core.AddToMapIfNotDefaultValueFloat32";
                break;
            case "list":
                mem.type = "[]string";
                break;

            default:
                break;
        }
    });
}
