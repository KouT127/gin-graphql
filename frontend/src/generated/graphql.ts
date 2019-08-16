import gql from 'graphql-tag';
export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
  Time: any,
};

export type AddTaskPayload = {
  __typename?: 'AddTaskPayload',
  clientMutationId?: Maybe<Scalars['String']>,
  task?: Maybe<Task>,
};

export type AddUserPayload = {
  __typename?: 'AddUserPayload',
  clientMutationId?: Maybe<Scalars['String']>,
  user?: Maybe<User>,
};

export type Connection = {
  pageInfo?: Maybe<PageInfo>,
  edges: Array<Edge>,
};

export type Edge = {
  cursor?: Maybe<Scalars['String']>,
};

export type Mutation = {
  __typename?: 'Mutation',
  addUser: AddUserPayload,
  addTask: AddTaskPayload,
};


export type MutationAddUserArgs = {
  user: UserInput
};


export type MutationAddTaskArgs = {
  input: TaskInput
};

export type Node = {
  id: Scalars['ID'],
};

export type PageInfo = {
  __typename?: 'PageInfo',
  startCursor?: Maybe<Scalars['String']>,
  endCursor?: Maybe<Scalars['String']>,
  hasNextPage: Scalars['Boolean'],
  hasPreviousPage: Scalars['Boolean'],
};

export type Query = {
  __typename?: 'Query',
  user?: Maybe<User>,
  tasks: TaskConnection,
};


export type QueryUserArgs = {
  id?: Maybe<Scalars['Int']>
};


export type QueryTasksArgs = {
  first?: Maybe<Scalars['Int']>,
  after?: Maybe<Scalars['String']>,
  last?: Maybe<Scalars['Int']>,
  before?: Maybe<Scalars['String']>,
  query?: Maybe<Scalars['String']>
};

export type Task = {
  __typename?: 'Task',
  id: Scalars['ID'],
  title: Scalars['String'],
  description: Scalars['String'],
  user?: Maybe<User>,
};

export type TaskConnection = {
  __typename?: 'TaskConnection',
  totalCount: Scalars['Int'],
  edges: Array<TaskEdge>,
  pageInfo?: Maybe<PageInfo>,
};

export type TaskEdge = {
  __typename?: 'TaskEdge',
  cursor: Scalars['String'],
  node: Task,
};

export type TaskInput = {
  title: Scalars['String'],
  description: Scalars['String'],
  userId: Scalars['String'],
};


export type User = {
  __typename?: 'User',
  id: Scalars['ID'],
  name: Scalars['String'],
  gender: Scalars['String'],
  tasks?: Maybe<TaskConnection>,
};


export type UserTasksArgs = {
  first?: Maybe<Scalars['Int']>,
  after?: Maybe<Scalars['String']>,
  last?: Maybe<Scalars['Int']>,
  before?: Maybe<Scalars['String']>,
  query?: Maybe<Scalars['String']>
};

export type UserConnection = {
  __typename?: 'UserConnection',
  totalCount: Scalars['Int'],
  edges: Array<UserEdge>,
  pageInfo?: Maybe<PageInfo>,
};

export type UserEdge = {
  __typename?: 'UserEdge',
  cursor: Scalars['String'],
  node: User,
};

export type UserInput = {
  name: Scalars['String'],
  gender: Scalars['String'],
  tasks?: Maybe<Array<Maybe<TaskInput>>>,
};
