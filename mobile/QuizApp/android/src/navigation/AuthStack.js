import React from 'react';
import {createStackNavigator} from '@react-navigation/stack';
import LoginPage from '../screens/Auth/LoginPage';
import RegistrationPage from '../screens/Auth/RegistrationPage';

const Stack = createStackNavigator();

const AuthStack = () => {
  return (
    <Stack.Navigator initialRouteName="Login">
      <Stack.Screen name="Login" component={LoginPage} />
      <Stack.Screen name="Registration" component={RegistrationPage} />
    </Stack.Navigator>
  );
};

export default AuthStack;
