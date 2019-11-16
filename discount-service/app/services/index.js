const discountService = require('./discount.service')

module.exports = [
  {
    protoPath: 'proto/discount.proto',
    name: 'DiscountService',
    service: discountService
  }
]