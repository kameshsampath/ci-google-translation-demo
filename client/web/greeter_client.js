const {GreeterClient} = require("./greeter_grpc_web_pb")
require("./greeter_pb")

var greeterService = new GreeterClient("http://localhost:10000")

var req = new proto.greeter.TranslationRequest()
req.setMessage("Hello World!")
req.setSourcelang("en")
req.addTargetlangs("ta")
req.addTargetlangs("te")
req.addTargetlangs("kn")

var stream = greeterService.greet(req)

stream.on("data",function (res) {
  console.trace("res:%s",JSON.stringify(res))
  document.getElementById("tmessage").innerHTML=res.array[0]
})

stream.on("error",function (err) {
  console.error(`Unexpected stream error: code = ${err.code}` +
              `, message = "${err.message}"`);
})