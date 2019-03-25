const axios = require ('axios')

module.exports = class{
    constructor(settings){
        console.log(settings)
        if(!settings.url){
            throw new Error('url must be defined')
        }

        this.settings = settings;
    }
    newEvent(inputs){
        if(!inputs.name){
            console.log('name must be defined')
            return
        }
        if(!inputs.user_id){
            console.log('user_id must be defined')
            return
        }
        const url = this.settings.url + '/event'
        axios.post(url, inputs)
        .then(function (response){
            console.log(response);
        })
        .catch(function (error){
            console.log(error)
        })
    }
    newUser(inputs){
        if(!inputs.username){
            console.log('username must be defined')
            return
        }
        if(!inputs.email){
            console.log('email must be defined')
            return
        }
        const url = this.settings.url + '/url'
        axios.post(url, inputs)
        .then(function (response){
            console.log(response);
        })
        .catch(function (error){
            console.log(error)
        })
    }
}