const Analytics = require ('../sdk-js')


const sdk = new Analytics({
    url: 'http://localhost:1337'
})

sdk.newEvent({
    name: 'L evenement est incroyable',
    user_id: '1IMK4J67IzRSPYMDkFIC5ep1pwb'
})


sdk.newUser({
    username: 'Dibouli',
    email: 'hey@hey.fr'
})
