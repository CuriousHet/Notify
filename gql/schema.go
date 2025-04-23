package gql

const SchemaString = `
    type Query {
        notifications(UserId: String!): [String!]!
    }
`
