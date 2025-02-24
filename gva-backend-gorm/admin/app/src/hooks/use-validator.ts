import type { Rule } from 'ant-design-vue/es/form';

import { $t } from '@vben/locales';

export type FormRule = Record<string, Rule[]>;

export const useValidator = () => {
  const checkEnum = (enums: string[], message?: string): Rule => {
    return {
      enum: enums,
      message,
    };
  };

  const required = (field?: string, message?: string): Rule => {
    return {
      message:
        message ||
        (field && $t('validator.error.required', { field })) ||
        $t('validator.error.noEmpty'),
      required: true,
    };
  };

  const lengthRange = (min: number, max: number, message?: string): Rule => {
    return {
      max,
      message: message || $t('validator.error.lengthRange', { max, min }),
      min,
    };
  };

  const noUnusualSpace = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        const noLeadingSpace = /^\s/;
        if (noLeadingSpace.test(val)) {
          callback(message || $t('validator.error.noLeadingSpace'));
        }
        const noMultipleSpaces = /\s{2,}/;
        if (noMultipleSpaces.test(val)) {
          callback(message || $t('validator.error.noMultipleSpaces'));
        }
        callback();
      },
    };
  };

  const noSpace = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        if (val?.indexOf(' ') === -1) {
          return Promise.resolve();
        } else {
          callback(message || $t('validator.error.noSpace'));
        }
      },
    };
  };

  const noNumber = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        if (val && /\d/.test(val)) {
          callback(message || $t('validator.error.noNumber'));
        } else {
          callback();
        }
      },
    };
  };

  const noSpecialCharacters = (message?: string): Rule => {
    const regexrule = /[!"#$%&'()*+,./:;<>?@[\\\]^_`{}~]/g;

    return {
      validator: (_, val, callback) => {
        if (regexrule.test(val)) {
          const matches = val.match(regexrule);
          let specialCharactors = '';
          if (matches) {
            specialCharactors += '`';
            specialCharactors += matches.join(', '); // Join the matched characters into a string
            specialCharactors += '`';
          }

          callback(
            message ||
              $t('validator.error.noSpecialCharacters', { specialCharactors }),
          );
        } else {
          callback();
        }
      },
    };
  };

  const defaultOpts = {
    allowNumber: false,
    allowSpecialCharacters: false,
    allowUnusualSpace: false,
  };

  const text = (
    opts: {
      allowNumber?: boolean;
      allowSpecialCharacters?: boolean;
      allowUnusualSpace?: boolean;
    } = defaultOpts,
  ): Rule => {
    return {
      validator: (_, val, callback) => {
        if (val) {
          const _callback = (error?: string) => {
            if (error) {
              callback(error);
            }
          };
          if (!opts.allowNumber) {
            noNumber().validator?.(_, val, _callback);
          }

          if (!opts.allowSpecialCharacters) {
            noSpecialCharacters().validator?.(_, val, _callback);
          }

          if (!opts.allowUnusualSpace) {
            noUnusualSpace().validator?.(_, val, _callback);
          }
        }
        callback();
      },
    };
  };

  const phone = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        if (!val) return callback();
        if (/^1[3-9]\d{9}$/.test(val)) {
          callback();
        } else {
          callback(message || $t('validator.error.phone'));
        }
      },
    };
  };

  const email = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        if (!val) return callback();
        const emailRegex = /^(?:\w-*\.*)+@(?:\w-?)+(?:\.\w{2,})+$/;
        if (emailRegex.test(val)) {
          callback();
        } else {
          callback(message || $t('validator.error.email'));
        }
      },
    };
  };

  const maxlength = (max: number, message?: string): Rule => {
    return {
      max,
      message: message || $t('validator.error.maxLength', { max }),
    };
  };

  const check = (message?: string): Rule => {
    return {
      validator: (_, val, callback) => {
        if (val) {
          callback();
        } else {
          callback(message || $t('validator.error.required'));
        }
      },
    };
  };

  const pattern = (pattern: RegExp, message?: string): Rule => {
    return {
      pattern,
      message: message || $t('validator.error.pattern'),
    };
  };

  const ip = (message?: string): Rule => {
    return pattern(/^((25[0-5]|(2[0-4]|1\d|[1-9]?)\d)\.?){4}$/, message);
  };

  const isNumber = (message?: string): Rule => {
    return {
      pattern: /^\d*$/,
      message: message || $t('validator.error.number'),
    };
  };

  const minMaxNumber = (min: number, max: number, message?: string): Rule => {
    return {
      message: message || $t('validator.error.minMaxNumber', { max, min }),
      validator: (_, val, callback) => {
        if (val >= min && val <= max) {
          callback();
        } else {
          callback(message || $t('validator.error.minMaxNumber', { max, min }));
        }
      },
    };
  };

  return {
    check,
    checkEnum,
    email,
    lengthRange,
    maxlength,
    text,
    notSpace: noSpace,
    notSpecialCharacters: noSpecialCharacters,
    phone,
    pattern,
    minMaxNumber,
    required,
    ip,
    isNumber,
    noUnusualSpace,
  };
};
