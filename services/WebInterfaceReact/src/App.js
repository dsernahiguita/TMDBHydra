import React from 'react';
import { Route } from 'react-router-dom';
import { ThemeProvider } from '@material-ui/styles';
import { hydraTheme } from './theme/hydra.theme';
import Home from './pages/home'

function App() {
  return (
    <ThemeProvider theme={hydraTheme}>
      <main>
        <Route exact path='/' render={(props) => <Home {...props} />} />
      </main>
    </ThemeProvider>
  );
}

export default App;
