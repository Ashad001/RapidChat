import React from 'react';
import auth from './auth';
import { Route, Redirect } from 'react-router-dom';

export const SecureRoute = ({component: Component, ...rest}) => {
  // Check if the user is authenticated
  if (auth.isAuthenticated()) {
    // If authenticated, render the provided component
    return (
      <Route
        {...rest}
        render={(props) => <Component {...props} />}
      />
    );
  } else {
    // If not authenticated, redirect to the home page
    return (
      <Route
        {...rest}
        render={(props) => <Redirect to={{pathname: '/', state: {from: props.location}}} />}
      />
    );
  }
};
