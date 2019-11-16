const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const { map, reduce } = require('ramda')

const loadProto = (path, service) => {
  const def = grpc.loadPackageDefinition(protoLoader.loadSync(path, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }))

  return def[service].service
}

const loadServiceProtos =
  map(s => ({ def: loadProto(s.protoPath, s.name), impl: s.service }))

const loadServices = services => server =>
  reduce(
    (server, s) => {
      server.addService(s.def, s.impl)
      return server
    }, server, loadServiceProtos(services)
  )

module.exports = { loadServices }