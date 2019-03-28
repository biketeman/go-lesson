const axios = require('axios');

module.exports = class {
    constructor(settings) {

        if (!settings.url) {
            settings.url = 'http://'
        }

        this.settings = settings;
    }

    newEvent(inputs) {

        if (!inputs.name) {
            console.log('name must be defined')
            return
        }

        if (!inputs.user_id) {
            console.log('user_id must be defined')
            return
        }

        const url = this.settings.url + '/event'
        axios.post(url, inputs)
          .then(function (response) {
            console.log(response.data);
          })
          .catch(function (error) {
            // console.log(error);
          });
    }

    newUser(inputs) {

        if (!inputs.username) {
            console.log('username must be defined');
            return;
        }

        if (!inputs.email) {
            console.log('email must be defined');
            return;
        }

        const url = this.settings.url + '/user'
        axios.post(url, inputs)
          .then(function (response) {
            console.log(response.data);
          })
          .catch(function (error) {
            // console.log(error);
          });
    }
}