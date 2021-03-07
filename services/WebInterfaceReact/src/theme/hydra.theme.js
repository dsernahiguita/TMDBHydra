import { createMuiTheme } from '@material-ui/core/styles';

export const hydraTheme = createMuiTheme({
  typography: {
    h5: {
      color: '#472C68'
    },
    subtitle1: {

    },
    button: {
      textTransform: 'none',
    },
    textPrimary: {
      color: '#472C68'
    }

  },
  overrides: {
    MuiOutlinedInput: {
      root: {
        backgroundColor: 'rgba(255,255,255,1)',
      }
    }
  },
  shape:{},
  palette: {
    primary: {
      // light: will be calculated from palette.primary.main,
      main: '#472C68'
      // dark: will be calculated from palette.primary.main,
      // contrastText: will be calculated to contrast with palette.primary.main
    },
    secondary: {
      // light: will be calculated from palette.secondary.main,
      main: '#00AABB'
      // dark: will be calculated from palette.secondary.main,
      // contrastText: will be calculated to contrast with palette.secondary.main
    },
    // Used by `getContrastText()` to maximize the contrast between
    // the background and the text.
    // contrastThreshold: 3,
    // Used by the functions below to shift a color's luminance by approximately
    // two indexes within its tonal palette.
    // E.g., shift from Red 500 to Red 300 or Red 700.
    // tonalOffset: 0.2,
  },
});
