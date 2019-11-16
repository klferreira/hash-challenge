const User = require('../models/User')
const Product = require('../models/Product')

var BLACK_FRIDAY = {
  day: 25,
  month: 11,
}

function compareDayAndMonth(d1, d2) {
  return d1.getUTCDate() == d2.getUTCDate() && d2.getUTCMonth() == d2.getUTCMonth()
}

function isBlackFriday() {
  const today = new Date()
  const blackFriday = new Date(`${today.getFullYear()}-${BLACK_FRIDAY.month}-${BLACK_FRIDAY.day}`)
  return compareDayAndMonth(today, blackFriday)
}

function isUserBirthday(birthday) {
  const today = new Date()
  return compareDayAndMonth(today, birthday)
}

async function getProductDiscount({ productID, userID }) {
  let pct = 0

  const dob = await User.findByPk(userID)
    .then(user => new Date(user.get('dateOfBirth')))
    .catch(err => Promise.reject(err))

  if (isUserBirthday(dob)) 
    pct = .05 

  if (isBlackFriday())
    pct = .10 

  const priceInCents = await Product.findByPk(productID)
    .then(p => p.priceInCents)
    .catch(err => Promise.reject(err))

  return { 
    discount: {
      priceInCents: priceInCents * pct, 
      percentual: pct * 100
    }
  }
}

module.exports = {
  GetDiscount: async (call, callback) => {
    try {
      callback(null, await getProductDiscount(call.request))
    } catch(err) {
      callback(null, err)
    }
  }
}