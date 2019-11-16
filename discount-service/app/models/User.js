const Sequelize = require('sequelize')
const db = require('../../config/database')

const User = db.define('user', {
  id: {
    primaryKey: true,
    type: Sequelize.STRING,
  },
  firstName: {
    field: 'first_name',
    type: Sequelize.STRING,
  },
  lastName: {
    field: 'last_name',
    type: Sequelize.STRING,
  },
  dateOfBirth: {
    field: 'date_of_birth',
    type: Sequelize.STRING,
  }
}, {
  timestamps: false,
})

module.exports = User