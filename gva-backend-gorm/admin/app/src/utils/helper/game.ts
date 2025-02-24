import { getLunar } from 'chinese-lunar-calendar';

import {
  GAME_TYPE,
  LIU_HE_BLUE_NUMS,
  LIU_HE_RED_NUMS,
  LUN_PAN_BLACK_ROULETTE_NUMS,
  LUN_PAN_RED_ROULETTE_NUMS,
  ZODIAC_SIGN,
} from '#/constants/game';

import { formatToDate } from './date-util';

export const analysisDict = {
  equal: {
    label: '和',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
  small: {
    label: '小',
    textColor: 'text-[#F28C28]',
    backgroundColor: 'bg-[#F28C28]',
  },
  big: {
    label: '大',
    textColor: 'text-blue-500',
    backgroundColor: 'bg-blue-500',
  },
  even: {
    label: '双',
    textColor: 'text-[#F28C28]',
    backgroundColor: 'bg-[#F28C28]',
  },
  odd: {
    label: '单',
    textColor: 'text-blue-500',
    backgroundColor: 'bg-blue-500',
  },
  red: {
    label: '红',
    textColor: 'text-[#ED2A23]',
    backgroundColor: 'bg-[#ED2A23] text-white',
  },
  green: {
    label: '绿',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
  smallZero: {
    label: '零',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
  blue: {
    label: '蓝',
    textColor: 'text-[#1699DD]',
    backgroundColor: 'bg-[#1699DD]',
  },
  black: {
    label: '黑',
    textColor: 'text-black dark:text-white',
    backgroundColor: 'bg-black text-white',
  },
  tiger: {
    label: '虎',
    textColor: 'text-[#ED2A23]',
    backgroundColor: 'bg-[#ED2A23]',
  },
  dragon: {
    label: '龙',
    textColor: 'text-[#1699DD]',
    backgroundColor: 'bg-[#1699DD]',
  },
  tigerDragonEqual: {
    label: '和',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
  player: {
    label: '闲',
    textColor: 'text-[#1699DD]',
    backgroundColor: 'bg-[#1699DD]',
  },
  banker: {
    label: '庄',
    textColor: 'text-[#ED2A23]',
    backgroundColor: 'bg-[#ED2A23]',
  },
  luiHeDraw: {
    label: '和局',
    textColor: 'text-black dark:text-white',
    backgroundColor: 'bg-black text-white',
  },
  leoPard: {
    label: '豹子',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
  fullCircumference: {
    label: '全围',
    textColor: 'text-[#16BE10]',
    backgroundColor: 'bg-[#16BE10] text-white',
  },
};

export const sumResult = (arr: number[]): number => {
  return arr.reduce((a, b) => a + b, 0);
};

export const zodiacs = (
  date: Date,
): { label: string; number: number; order: number; value: string }[] => {
  const [year, month, day] = formatToDate(date).split('-');

  /* Output
    lunarMonth: 12,   //农历月份
    lunarDate: 17,    //农历日期
    isLeap: false,    //是否闰月
    solarTerm: null,  //节气，null代表没有
    lunarYear: '庚午年', //农历年份，年以正月初一开始
    zodiac: '马',     //生肖，生肖以正月初一开始
    dateStr: '腊月十七' //农历中文
  */
  const currentLunarYear = getLunar(year, month, day);

  const startIndex = ZODIAC_SIGN.findIndex(
    (sign) => sign.label === currentLunarYear.zodiac,
  );

  const orderedZodiacSign = [
    ...ZODIAC_SIGN.slice(startIndex),
    ...ZODIAC_SIGN.slice(0, startIndex),
  ];

  return orderedZodiacSign.map((sign, index) => {
    return {
      value: sign.value,
      label: sign.label,
      number: index === 0 ? 1 : 13 - index,
      order: sign.order,
    };
  });
};

export const getZodiac = (num: number, date: Date, isZh = true): string => {
  let idx = num % 12;
  idx = idx === 0 ? 12 : idx;
  const zodiac = zodiacs(date).find((z) => z.number === idx);
  return (isZh ? zodiac?.label : zodiac?.value) || '';
};

export const getNextDrawIssue = (lastDrawIssue: number): string => {
  const nextDrawIssue = lastDrawIssue + 1 > 1440 ? 1 : lastDrawIssue + 1;
  let zeroStr = '';
  // the last draw issue length
  const loopTime = 4 - nextDrawIssue.toString().length;
  for (let i = 0; i < loopTime; i++) {
    zeroStr += '0';
  }
  return zeroStr + nextDrawIssue;
};

export const getBaiJiaLeNumber = (input: number): number => {
  if (input === 0) {
    return input;
  }
  const num = input % 13 || 13;
  const type =
    input % 13 === 0 ? Math.floor(input / 13) - 1 : Math.floor(input / 13);

  return num + 3 * (num - 1) + type;
};

export const isBaiJiaLePair = (arr: number[]): boolean => {
  // arr contains only 2 number
  const numOne = (arr[0] as number) % 13 || 13;
  const numTwo = (arr[1] as number) % 13 || 13;

  return numOne === numTwo;
};

export const isLongHuRed = (input: number): boolean => {
  const num =
    input % 13 === 0 ? Math.floor(input / 13) - 1 : Math.floor(input / 13);
  return num % 2 === 0;
};

// all cards have same number e.g: A,A,A  2,2,2   Q,Q,Q  K,K,K
export const isThreeCardAreSame = (arr: number[]): boolean => {
  return arr.every((item) => item % 13 === (arr[0] as number) % 13);
};

// all cards have same type e.g: all heart, all diamond
export const isFlush = (arr: number[]): boolean => {
  const firstNum =
    (arr[0] as number) % 13 === 0
      ? Math.floor((arr[0] as number) / 13) - 1
      : Math.floor((arr[0] as number) / 13);
  return arr.every((item) => {
    const num =
      item % 13 === 0 ? Math.floor(item / 13) - 1 : Math.floor(item / 13);
    return num === firstNum;
  });
};

// when the cards are straight / consecutive e.g: A,2,3   10,J,Q   J,Q,K
export const isStraight = (arr: number[]): boolean => {
  const straightArr: number[] = [];
  arr.forEach((a) => {
    const num = a % 13 === 0 ? 13 : a % 13;

    straightArr.push(num);
  });
  straightArr.sort();

  return straightArr.every((e, i) => {
    if (i === straightArr.length - 1) return true;
    return e !== undefined && e === (straightArr[i + 1] as number) - 1;
  });
};

export const isPair = (arr: number[]): boolean => {
  const firstCard = (arr[0] as number) % 13;
  const secondCard = (arr[1] as number) % 13;
  const thirdCard = (arr[2] as number) % 13;

  if (
    firstCard === secondCard ||
    firstCard === thirdCard ||
    secondCard === thirdCard
  )
    return true;

  return false;
};

export const isSanGongPairOrMore = (arr: number[]): boolean => {
  return (
    isPair(arr) || isThreeCardAreSame(arr) || isStraight(arr) || isFlush(arr)
  );
};

export const getStraightLine = (num: number): string => {
  if (num > 0 && num <= 12) {
    return '直1';
  } else if (num > 12 && num <= 24) {
    return '直2';
  }
  return '直3';
};

export const getCombination12Digit = (num: number): string => {
  if (num % 3 === 2) {
    return '组2';
  } else if (num % 3 === 1) {
    return '组1';
  } else {
    return '组3';
  }
};

export const getFinalGameTwoColor = (arr: number[]): string => {
  if (arr.length < 2) {
    return '';
  }

  let red = 0;
  let blue = 0;
  let green = 0;

  // Count occurrences of each color
  for (const element of arr) {
    if (LIU_HE_RED_NUMS.includes(element)) {
      red++;
    } else if (LIU_HE_BLUE_NUMS.includes(element)) {
      blue++;
    } else {
      green++;
    }
  }

  const lastElement = arr[arr.length - 1] as number;
  const isLastRed = LIU_HE_RED_NUMS.includes(lastElement);
  const isLastBlue = LIU_HE_BLUE_NUMS.includes(lastElement);

  if (red > blue) {
    const redDominates = red > green || (red === green && isLastRed);
    return redDominates ? '红' : '绿';
  } else if (red === blue && isLastRed) {
    return '红';
  } else {
    const blueDominates = blue > green || (blue === green && isLastBlue);
    return blueDominates ? '蓝' : '绿';
  }
};

export const isLeoPard = (arr: number[], is1n6: boolean): boolean => {
  if (!is1n6) {
    if (arr.every((element) => element === arr[0])) {
      return true;
    }
    return false;
  }

  if (
    (arr[0] === 1 || arr[0] === 6) &&
    arr.every((element) => element === arr[0])
  ) {
    return true;
  }

  return false;
};

export const getLeoPardColor = (
  type: 'backgroundColor' | 'label' | 'textColor',
): string => {
  return analysisDict.leoPard[type];
};

export const getSumLabelByGameType = (
  result: number[],
  gameType: GAME_TYPE,
): number => {
  switch (gameType) {
    case GAME_TYPE.HAPPY_EIGHT: {
      return sumResult(result.slice(0, 20));
    }
    case GAME_TYPE.KUAI_CHE: {
      return (result[0] as number) + (result[1] as number);
    }
    case GAME_TYPE.LIU_HE: {
      return result[result.length - 1] as number;
    }
    default: {
      return sumResult(result);
    }
  }
};
export const getBigSmallEqual = (props: {
  bigNumber: number;
  equalNumber?: number;
  num: number;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, num, bigNumber, equalNumber } = props;
  if (equalNumber && num === equalNumber) return analysisDict.equal[type];

  return num < bigNumber ? analysisDict.small[type] : analysisDict.big[type];
};

export const getBigSmallByGameType = (props: {
  gameType: string;
  result: number[];
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { gameType, result, type } = props;
  switch (gameType) {
    case GAME_TYPE.ELEVEN_PICK_FIVE: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 31,
        equalNumber: 30,
      });
    }
    case GAME_TYPE.HAPPY_EIGHT: {
      return getBigSmallEqual({
        type,
        num: sumResult(result.slice(0, 20)),
        bigNumber: 811,
        equalNumber: 810,
      });
    }
    case GAME_TYPE.HAPPY_TEN: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 84,
        equalNumber: 84,
      });
    }
    case GAME_TYPE.KUAI_CHE: {
      return getBigSmallEqual({
        type,
        num: (result[0] as number) + (result[1] as number),
        bigNumber: 12,
      });
    }
    case GAME_TYPE.KUAI_SHAN:
    case GAME_TYPE.YU_XIA_XIE: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 11,
      });
    }
    case GAME_TYPE.LIU_HE: {
      return getBigSmallEqual({
        type,
        num: result[result.length - 1] as number,
        bigNumber: 25,
      });
    }
    case GAME_TYPE.LUN_PAN: {
      return getBigSmallEqual({
        type,
        num: result[0] as number,
        bigNumber: 19,
      });
    }
    case GAME_TYPE.PC28: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 14,
      });
    }
    case GAME_TYPE.SHI_SHI: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 23,
      });
    }

    default: {
      return getBigSmallEqual({
        type,
        num: sumResult(result),
        bigNumber: 11,
      });
    }
  }
};

export const getOddEvenEqual = (props: {
  equalNumber?: number;
  num: number;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, num, equalNumber } = props;
  if (equalNumber && num === equalNumber) return analysisDict.equal[type];

  return num % 2 === 0 ? analysisDict.even[type] : analysisDict.odd[type];
};

export const getOddEvenLabelByGameType = (
  type: 'backgroundColor' | 'label' | 'textColor',
  result: number[],
  gameType: string,
) => {
  switch (gameType) {
    case GAME_TYPE.KUAI_CHE: {
      return getOddEvenEqual({
        type,
        num: (result[0] as number) + (result[1] as number),
      });
    }
    case GAME_TYPE.LIU_HE: {
      return getOddEvenEqual({
        type,
        num: result[result.length - 1] as number,
      });
    }
    default: {
      return getOddEvenEqual({ type, num: sumResult(result) });
    }
  }
};

export const isFullCircumference = (result: number[]) => {
  if (result.every((element) => element === result[0])) {
    return true;
  }
  return false;
};

export const getFullCircumference = (
  type: 'backgroundColor' | 'label' | 'textColor',
) => {
  return analysisDict.fullCircumference[type];
};

export const getTigerDragon = (props: {
  firstNum: number;
  secondNum: number;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, firstNum, secondNum } = props;
  if (firstNum === secondNum) return analysisDict.tigerDragonEqual[type];
  if (firstNum > secondNum) return analysisDict.dragon[type];
  return analysisDict.tiger[type];
};

export const getFirstVsLast = (
  type: 'backgroundColor' | 'label' | 'textColor',
  result: number[],
  gameType: string,
) => {
  const firstNum = result[0] as number;
  const lastNum = result[result.length - 1] as number;

  if (gameType === 'niu_niu') {
    return result[result.length - 1]
      ? getTigerDragon({
          type,
          firstNum: result[0] as number,
          secondNum: result[4] as number,
        })
      : getTigerDragon({
          type,
          firstNum: result[6] as number,
          secondNum: result[10] as number,
        });
  }

  return getTigerDragon({ type, firstNum, secondNum: lastNum });
};

export const getRedGreenBlue = (props: {
  num: number;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, num } = props;
  if (LIU_HE_RED_NUMS.includes(num)) return analysisDict.red[type];
  else if (LIU_HE_BLUE_NUMS.includes(num)) return analysisDict.blue[type];
  return analysisDict.green[type];
};

export const getRedBlackGreen = (props: {
  result: number | undefined;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, result } = props;
  if (LUN_PAN_RED_ROULETTE_NUMS.includes(result as number))
    return analysisDict.red[type];
  else if (LUN_PAN_BLACK_ROULETTE_NUMS.includes(result as number))
    return analysisDict.black[type];
  return analysisDict.smallZero[type];
};

export const getLongHu = (props: {
  num: number | undefined;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, num } = props;
  if (num === 0) return analysisDict.dragon[type];
  if (num === 1) return analysisDict.tiger[type];
  return analysisDict.equal[type];
};

export const getLongHuColor = (props: {
  num: number | undefined;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { num, type } = props;

  return isLongHuRed(num as number)
    ? analysisDict.red[type]
    : analysisDict.black[type];
};

export const getBankerPlayer = (props: {
  num: number | undefined;
  type: 'backgroundColor' | 'label' | 'textColor';
}) => {
  const { type, num } = props;

  if (num === 0) return analysisDict.banker[type];
  if (num === 1) return analysisDict.player[type];
  return analysisDict.equal[type];
};
