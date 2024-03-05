import React from 'react';
import {View, Text, StyleSheet} from 'react-native';
import Button from '../../components/Button';

const LoginPage = () => {
  const handleLogin = () => {
    console.log('User logged in');
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Login Page</Text>

      <Button onPress={handleLogin} title="Login" />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
  },
});

export default LoginPage;
