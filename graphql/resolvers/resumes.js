const {ObjectID} = require('mongodb');

const resolver = async (parent, args, context, info) => {
  const {first = 20, after} = args;
  const {db} = context;
  const collection = db.collection("resumes");
  const totalCount = collection.countDocuments();
  const resumes = await collection.find(after ? {_id: {$gt: ObjectID(after)}} : {}).limit(first).toArray();
  const lastResume = (await collection.find({}).skip(totalCount - 1).limit(1).toArray())[0];
  const edges = resumes.map((resume) => ({node: resume}));
  const ids = resumes.map((resume) => resume._id.toString());
  const endCursor = resumes[resumes.length - 1]._id;

  return {
    totalCount,
    edges,
    pageInfo: {
      hasNextPage: !ids.includes(lastResume._id.toString()),
      endCursor,
    },
  };
};

module.exports = resolver;
