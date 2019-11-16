const Sequelize = require('sequelize')

function getDatabaseURI(uri = 'root:toor@127.0.0.1:3306/hashchallenge') {
	return `mysql://${uri}`
}

module.exports = new Sequelize(getDatabaseURI(process.env.DATABASE_URI))