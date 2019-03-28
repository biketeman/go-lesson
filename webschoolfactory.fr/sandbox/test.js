const Analytics = require('../sdk-js');

const sdk = new Analytics({
    url: 'http://localhost:1337'
})

// sdk.newUser({

// })

sdk.newEvent({
    // name: 'Mon super event',
    user_id: '1IwsZTolFQM12Ynen248xrM8yFL'
})