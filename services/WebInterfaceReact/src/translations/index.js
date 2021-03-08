import React from 'react';
import Polyglot from 'node-polyglot';

const englishTranslation = require('./en');
const germanTranslation = require('./de');
const spanishTranslation = require('./es');

export const LANGUAGE_EN = 'en';
export const LANGUAGE_DE = 'de';
export const LANGUAGE_ES = 'es';

export const DEFAULT_LANGUAGE = 'es';
let language;

const languages = {
  en: { name: 'English', translation: englishTranslation },
  es: { name: 'Español', translation: spanishTranslation },
  de: { name: 'Deutsch', translation: germanTranslation },
};

export const setLanguage = (newLanguage) => {
  language = newLanguage;
};


export const withPolyglot = (WrappedComponent) => {
  /* Wrapped component could not extends for Pure.Component, please don't change it */
  class PolyglotClass extends React.Component {
    render() {
      if (!language) {
        language = DEFAULT_LANGUAGE;
      }
      const polyglot = new Polyglot({
        phrases: languages[language].translation,
        locale: language,
        currentLocale: language,
      });
      return <WrappedComponent {...this.props} polyglot={polyglot} />;
    }
  }
  return PolyglotClass;
};


export default withPolyglot;
