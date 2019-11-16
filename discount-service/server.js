const server = require('./config/grpc')
const db = require('./config/database')

const services = require('./app/services')

db.authenticate()
  .then(() => console.log("Connected to the database"))
  .catch(err => console.log(`Failed to connect to database: ${err}`))

server.main(services, process.env.DISCOUNT_SERVICE_PORT)