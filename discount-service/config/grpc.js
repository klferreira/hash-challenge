const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const { map } = require('ramda')

const loadProto = (path, service) => {
  const def = grpc.loadPackageDefinition(protoLoader.loadSync(path, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }))

  return def.proto[service].service
}

const loadServiceProtos =
  map(s => ({ def: loadProto(s.protoPath, s.name), impl: s.service }))

function main(services, port = ":50052") {
  const server = new grpc.Server()
  
  loadServiceProtos(services)
    .forEach(service => server.addService(service.def, service.impl))

  server.bind(`0.0.0.0${port}`, grpc.ServerCredentials.createInsecure())

  server.start()
}

module.exports = { main }