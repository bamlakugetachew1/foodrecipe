import gql from "graphql-tag";

export var ADD_COMMENTS_MUTATION = gql`
  mutation MyMutation2($recipe_id: Int, $user_id: Int, $comment_text: String) {
    insert_comments(
      objects: {
        recipe_id: $recipe_id
        user_id: $user_id
        comment_text: $comment_text
      }
    ) {
      affected_rows
    }
  }
`;

export var LIKE_RECIPE_MUTATION = gql`
  mutation MyMutation($recipe_id: Int, $user_id: Int) {
    insert_likes(objects: { recipe_id: $recipe_id, user_id: $user_id }) {
      returning {
        id
        recipe_id
        user_id
      }
    }
  }
`;

export var UNLIKE_RECIPE_MUTATION = gql`
  mutation MyMutation($recipe_id: Int, $user_id: Int) {
    delete_likes(
      where: { recipe_id: { _eq: $recipe_id }, user_id: { _eq: $user_id } }
    ) {
      affected_rows
    }
  }
`;

export var ADD_RATING_MUTATION = gql`
  mutation MyMutation($recipe_id: Int, $user_id: Int, $rating: Int) {
    insert_ratings(
      objects: { recipe_id: $recipe_id, user_id: $user_id, rating: $rating }
    ) {
      returning {
        id
        rating
        recipe_id
        user_id
      }
    }
  }
`;

export var REMOVE_RATING_MUTATION = gql`
  mutation MyMutation($recipe_id: Int, $user_id: Int) {
    delete_ratings(
      where: { recipe_id: { _eq: $recipe_id }, user_id: { _eq: $user_id } }
    ) {
      returning {
        id
        rating
        recipe_id
        user_id
      }
    }
  }
`;
