export type NextErrorProps = {
  error: Error & { digest?: string };
  reset: () => void;
};
