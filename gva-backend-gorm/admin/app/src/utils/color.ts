interface LevelColor {
  minLevel: number;
  maxLevel: number;
  color: string;
}

const levelColors: LevelColor[] = [
  { minLevel: 0, maxLevel: 9, color: 'border-[#C97755]' },
  { minLevel: 10, maxLevel: 19, color: 'border-[#76919E]' },
  { minLevel: 20, maxLevel: 29, color: 'border-[#E3A300]' },
  { minLevel: 30, maxLevel: 39, color: 'border-[#FE6B21]' },
  { minLevel: 40, maxLevel: 49, color: 'border-[#DA3D0D]' },
  { minLevel: 50, maxLevel: 59, color: 'border-[#FF6978]' },
  { minLevel: 60, maxLevel: 69, color: 'border-[#C50FF3]' },
  { minLevel: 70, maxLevel: 79, color: 'border-[#C515A4]' },
  { minLevel: 80, maxLevel: 89, color: 'border-[#801DB2]' },
  { minLevel: 90, maxLevel: Infinity, color: 'border-[#F42020]' },
];

export const getUserLevelBorderColor = (userLevel: number): string => {
  const levelColor = levelColors.find(
    ({ minLevel, maxLevel }) => userLevel >= minLevel && userLevel <= maxLevel,
  );
  return levelColor ? levelColor.color : '';
};
