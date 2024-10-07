import { ApolloClient, InMemoryCache, HttpLink } from "@apollo/client/core";
import { setContext } from "@apollo/client/link/context";
import { provideApolloClient } from "@vue/apollo-composable";
import { useUserStore } from "@/store/userStore"; // Import the user store

const httpLink = new HttpLink({
  uri: "http://localhost:8080/v1/graphql",
});

// Set the authentication context (headers)
const authLink = setContext((_, { headers, ...context }) => {
  const store = useUserStore(); // Access the user store
  const token = store.token; // Get the token from the store
  const isPublicQuery = context.isPublicQuery || false;
  if (isPublicQuery) {
    return {
      headers: {
        ...headers,
      },
    };
  }

  return {
    headers: {
      ...headers,
      Authorization: token ? `Bearer ${token}` : "",
    },
  };
});

// Create the Apollo Client instance with authLink and httpLink
const apolloClient = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default defineNuxtPlugin(() => {
  provideApolloClient(apolloClient);
});
