import React from 'react'
import auth from './auth'
import { Route, Redirect } from 'react-router-dom'

export const ProtectedRoute = ({component: Component, ...rest}) => {

  // return a route with component passed in
  // check if user is logged in
  return (
    <Route
      {...rest}
      render={(props) => auth.isAuthenticated()
        ? <Component {...props} />
        : <Redirect to={{pathname: '/', state: {from: props.location}}} />}
    />
  )
}