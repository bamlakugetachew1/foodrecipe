import gql from "graphql-tag";
export const GET_MY_RECIPES = gql`
  query MyQuery {
    recipes(order_by: { created_at: desc }) {
      category
      description
      id
      preparation_time
      title
      featured_images {
        image_url
      }
    }
  }
`;
