//Create tabs for “Shipping” and “Details” that display the shipping cost and product details, respectively.

var eventBus = new Vue()

var app = new Vue({
    el: '#app',
    data: {
        premium: true,
        cart: []
    },
    methods: {
        updateCart(id) {
            this.cart.push(id)
        }
    }
})