import { LanguageFaMain } from '@/languages/fa/topics/dashboard/main.language';
import { LanguageFaComponents } from './topics/common/components.language';
import { LanguageFaError } from './topics/common/error.language';
import { LanguageFaForm } from './topics/common/form.language';
import { LanguageFaHTTP } from './topics/common/http.language';
import { LanguageFaAuth } from './topics/public/auth.language';
import { LanguageFaRoot } from './topics/public/root.language';

export const LanguageFa = {
  publ: {
    root: LanguageFaRoot,
    auth: LanguageFaAuth,
  },
  dash: {
    main: LanguageFaMain,
  },
  comn: {
    erro: LanguageFaError,
    form: LanguageFaForm,
    http: LanguageFaHTTP,
    cmpn: LanguageFaComponents,
  },
};
