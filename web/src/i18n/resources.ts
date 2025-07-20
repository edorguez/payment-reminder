import en from '../locales/en/common.json';
import es from '../locales/es/common.json';

const resources = {
  en: {
    common: en
  },
  es: {
    common: es
  }
} as const;

export default resources;

// Type definitions
declare module 'i18next' {
  interface CustomTypeOptions {
    defaultNS: 'common';
    resources: typeof resources.en;
  }
}
