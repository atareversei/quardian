import { AnimatePresence, motion } from 'motion/react';
import { Group, MantineColor, Stack, Text } from '@mantine/core';
import { getColorCustomVariables } from '@/utils/color/color';
import { i } from '@/utils/icon/icon';
import c from './statements.module.css';

export type ValidationStatementStatus = 'failed' | 'neutral' | 'passed';

export type ValidationStatement = {
  text: string;
  passed: ValidationStatementStatus;
};

type Props = {
  statements: ValidationStatement[];
};

export function Statements({ statements }: Props) {
  function getIcon(status: ValidationStatementStatus): React.ReactNode {
    switch (status) {
      case 'failed':
        return i('x', 'sm', getColorCustomVariables('redText'));
      case 'neutral':
        return i('pointFilled', 'sm', getColorCustomVariables('dimmed'));
      case 'passed':
        return i('check', 'sm', getColorCustomVariables('greenText'));
    }
  }

  function getColor(status: ValidationStatementStatus): MantineColor {
    switch (status) {
      case 'failed':
        return 'red';
      case 'neutral':
        return 'dimmed';
      case 'passed':
        return 'green';
    }
  }

  function renderStatements() {
    return statements.map((statement) => (
      <Group gap={8} key={statement.text}>
        <div className={c.icon}>
          <AnimatePresence>
            <motion.div
              key={statement.passed}
              className={c.icon}
              initial={{ y: -20, opacity: 0.2, transition: { duration: 0.25 } }}
              animate={{ y: 0, opacity: 1 }}
              exit={{ opacity: 0.2, y: 20, transition: { duration: 0.15 } }}
            >
              {getIcon(statement.passed)}
            </motion.div>
          </AnimatePresence>
        </div>
        <Text size="sm" c={getColor(statement.passed)} className={c.text}>
          {statement.text}
        </Text>
      </Group>
    ));
  }

  return <Stack gap={2}>{renderStatements()}</Stack>;
}
