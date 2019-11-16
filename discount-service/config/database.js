const Sequelize = require('sequelize')

const { DB_USER, DB_PASS, DB_NAME, DB_URI } = process.env

const databaseUri = `mysql://${DB_USER}:${DB_PASS}@${DB_URI}/${DB_NAME}`

console.log(databaseUri)

module.exports = new Sequelize(databaseUri)