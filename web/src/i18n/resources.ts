import en from '../../public/locales/en/common.json';
import es from '../../public/locales/es/common.json';

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
