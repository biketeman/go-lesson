const { objectType } = require('nexus');

module.exports = objectType({
    name: 'User',
    description: '',
    definition(t) {
        t.string('name', {
            description: 'The name of the user.',
            nullable: false
        })
    }
});