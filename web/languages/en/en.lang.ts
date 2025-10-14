import { LanguageEnComponents } from './topics/common/components.language';
import { LanguageEnError } from './topics/common/error.language';
import { LanguageEnForm } from './topics/common/form.language';
import { LanguageEnHTTP } from './topics/common/http.language';
import { LanguageEnMain } from './topics/dashboard/main.language';
import { LanguageEnAuth } from './topics/public/auth.language';
import { LanguageEnRoot } from './topics/public/root.language';

export const LanguageEn = {
  publ: {
    root: LanguageEnRoot,
    auth: LanguageEnAuth,
  },
  dash: {
    main: LanguageEnMain,
  },
  comn: {
    erro: LanguageEnError,
    form: LanguageEnForm,
    http: LanguageEnHTTP,
    cmpn: LanguageEnComponents,
  },
};

