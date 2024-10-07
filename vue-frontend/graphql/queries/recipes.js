import gql from "graphql-tag";

export const CHECK_USER_RATED_RECIPE = gql`
  query MyQuery($recipe_id: Int, $user_id: Int) {
    ratings_aggregate(
      where: { recipe_id: { _eq: $recipe_id }, user_id: { _eq: $user_id } }
    ) {
      aggregate {
        count
      }
    }
  }
`;

export const GET_RECIPE_STEPS = gql`
  query MyQuery($recipe_id: Int) {
    steps(where: { recipe_id: { _eq: $recipe_id } }) {
      instruction
    }
  }
`;

export const CHECK_USER_LIKED_RECIPE = gql`
  query MyQuery($recipe_id: Int, $user_id: Int) {
    likes_aggregate(
      where: { recipe_id: { _eq: $recipe_id }, user_id: { _eq: $user_id } }
    ) {
      aggregate {
        count
      }
    }
  }
`;

export const GET_ALL_RECIPE_RATINGS = gql`
  query MyQuery($recipe_id: Int) {
    ratings_aggregate(where: { recipe_id: { _eq: $recipe_id } }) {
      aggregate {
        avg {
          rating
        }
        count
      }
    }
  }
`;

export const GET_ALL_RECIPE_COMMENTS = gql`
  query MyQuery($recipe_id: Int) {
    comments_aggregate(
      where: { recipe_id: { _eq: $recipe_id } }
      order_by: { created_at: desc }
    ) {
      aggregate {
        count
      }
      nodes {
        comment_text
        id
        recipe_id
        user_id
        created_at
      }
    }
  }
`;

export const GET_ALL_RECIPES_LIKES = gql`
  query MyQuery2($recipe_id: Int) {
    likes_aggregate(where: { recipe_id: { _eq: $recipe_id } }) {
      aggregate {
        count
      }
    }
  }
`;

export const GET_ALL_RECIPES_INGRIDENTS = gql`
  query MyQuery2($recipe_id: Int) {
    Ingredients(where: { recipe_id: { _eq: $recipe_id } }) {
      ingredient_name
    }
  }
`;

export const GET_ALL_RECIPES_IMAGES = gql`
  query MyQuery2($recipe_id: Int) {
    featured_images(where: { recipe_id: { _eq: $recipe_id } }) {
      image_url
      is_thumbnail
    }
  }
`;

export const RETRIVE_ALL_RECIPES = gql`
  query GetAllRecipes($limit: Int) {
    recipes(limit: $limit, order_by: { created_at: desc }) {
      category
      description
      preparation_time
      title
      id
      featured_images {
        image_url
      }
    }
  }
`;

export const FILTER_BY_TIME = gql`
  query myquery($limit: Int!) {
    recipes(limit: $limit, order_by: { preparation_time: asc }) {
      category
      description
      preparation_time
      title
      id
      featured_images {
        image_url
      }
    }
  }
`;

export const FILTER_BY_TITLE = gql`
  query myquery($limit: Int!) {
    recipes(limit: $limit, order_by: { title: desc }) {
      category
      description
      preparation_time
      title
      id
      featured_images {
        image_url
      }
    }
  }
`;

export const BROWSEBY_TITLE_CATEGORY_OR_CATAGORY = gql`
  query myquery($search: String!) {
    recipes(
      limit: 6
      where: {
        _or: [
          { title: { _like: $search } }
          { description: { _like: $search } }
          { category: { _like: $search } }
        ]
      }
    ) {
      category
      description
      preparation_time
      title
      id
      featured_images {
        image_url
      }
    }
  }
`;
