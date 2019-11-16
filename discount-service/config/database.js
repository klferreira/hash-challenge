const Sequelize = require('sequelize')

const { DB_USER, DB_PASS, DB_NAME, DB_URI } = process.env

function getDatabaseURI(dbUser = 'root', dbPass = 'toor', dbName = 'hashchallenge', uri = '127.0.0.1:3306') {
	return `mysql://${dbUser}:${dbPass}@${uri}/${dbName}`
}

module.exports = new Sequelize(getDatabaseURI(DB_USER, DB_PASS, DB_NAME, DB_URI))