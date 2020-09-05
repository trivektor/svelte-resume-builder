const resolver = async (parent, args, context) => {
    const {input: {name, description}} = args;
    const {db} = context;
    const collection = db.collection("resumes");

    const result = await collection.insertOne({name, description});

    return result.ops[0];
};

module.exports = resolver;