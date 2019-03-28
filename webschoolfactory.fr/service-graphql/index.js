const { queryType, stringArg, objectType, makeSchema } = require("nexus");
const { GraphQLServer } = require("graphql-yoga");

const User = require('./types/user')

const Event = objectType({
    name: 'Event',
    description: '',
    definition(t) {
        t.string('name', {
            description: 'The name of the event.',
            nullable: false
        }),
        t.string('created_at', {
            description: 'Timestamp of creation.',
            nullable: false
        })
    }
});

const Query = queryType({
    definition(t) {
        t.list.field("eventsByUserId", {
            type: Event,
            args: {
                user_id: stringArg({
                    nullable: false
                })
            },
            resolve: async (parent, args) => {

                const results = await knex.select('').from('events').where({
                    user_id: args.user_id
                })

                return results
            }
        });
    },
});

const schema = makeSchema({
    types: [Query, User, Event],
    outputs: {
        schema: __dirname + "/generated/schema.graphql",
        typegen: __dirname + "/generated/typings.ts",
    },
});

const server = new GraphQLServer({
    schema,
});

server.start({
    port: 1337
}, () => `Server is running on http://localhost:4000`);