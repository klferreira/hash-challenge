const Sequelize = require('sequelize')
const db = require('../../config/database')

const Product = db.define('product', {
  id: {
    primaryKey: true,
    type: Sequelize.STRING,
  },
  title: {
    type: Sequelize.STRING,
  },
  description: {
    type: Sequelize.STRING,
  },
  priceInCents: {
    field: 'price_in_cents',
    type: Sequelize.INTEGER,
  }
}, {
  timestamps: false,
})

module.exports = Product