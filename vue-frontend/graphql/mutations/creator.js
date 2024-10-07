import gql from "graphql-tag";

export var REGISTER_USER_MUTATION = gql`
  mutation registerUser($name: String!, $email: String!, $password: String!) {
    register(name: $name, email: $email, password: $password) {
      message
    }
  }
`;

export var LOGIN_USER_MUTATION = gql`
  mutation loginUser($email: String!, $password: String!) {
    loginUser(email: $email, password: $password) {
      id
      token
    }
  }
`;

export var INSERT_IMAGES_MUTATION = gql`
  mutation MyMutation($objects: [featured_images_insert_input!]!) {
    insert_featured_images(objects: $objects) {
      affected_rows
    }
  }
`;

export var INSERT_INGREDIENTS_MUTATION = gql`
  mutation MyMutation($objects: [Ingredients_insert_input!]!) {
    insert_Ingredients(objects: $objects) {
      affected_rows
    }
  }
`;

export var INSERT_STEPS_MUTATION = gql`
  mutation MyMutation($objects: [steps_insert_input!]!) {
    insert_steps(objects: $objects) {
      affected_rows
    }
  }
`;

export var DELETE_RECIPES_MUTATION = gql`
  mutation MyMutation($recipe_id: Int!) {
    delete_recipes_by_pk(id: $recipe_id) {
      category
      title
      preparation_time
      user_id
    }
  }
`;

export var INSERT_RECIPES_ONE_MUTAION = gql`
  mutation MyMutation2(
    $title: String
    $description: String
    $preparation_time: String
    $category: String
    $user_id: Int
  ) {
    insert_recipes_one(
      object: {
        title: $title
        description: $description
        preparation_time: $preparation_time
        category: $category
        user_id: $user_id
      }
    ) {
      id
    }
  }
`;
