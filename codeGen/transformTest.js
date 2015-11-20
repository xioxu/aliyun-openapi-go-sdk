var ejs = require('ejs');
var fs = require("fs")
var transoformerTem = fs.readFileSync(__dirname + '/transformer.ejs', 'utf8');

var props = [];

props.push({
	"IsArray":true,
	"tagName": "InstanceInfo",
	"basicType":"InstanceInfo",
	"typeName": "[]InstanceInfo",
	"Key":"InstanceInfo",
	"PropType":0
});

props.push({
	"IsArray":false,
	"tagName": "GroupInfo",
	"typeName": "GroupInfo",
	"Key":"GroupInfo",
	"PropType":2
});


props.push({
	"IsArray":false,
	"tagName": "RequestId",
	"typeName": "string",
	"Key":"RequestId",
	"GetValFunc":"GetString",
	"PropType":1
});

  var code = ejs.render(transoformerTem, {
        Props:props,
        ModalName: "AcsReqponse"
    });
	
console.log(code);

