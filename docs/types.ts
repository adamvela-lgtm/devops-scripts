/**
 * Represents the configuration for a DevOps script.
 */
export interface ScriptConfig {
  name: string;
  description: string;
  version: string;
  author: string;
  dependencies: Record<string, string>;
  environmentVariables: string[];
  timeout?: number;
}

/**
 * Represents the result of a script execution.
 */
export interface ScriptResult {
  success: boolean;
  output: string;
  error?: string;
  duration: number;
}

/**
 * Represents the supported operating systems for the scripts.
 */
export enum OperatingSystem {
  Linux = 'linux',
  Windows = 'windows',
  MacOS = 'darwin',
}

/**
 * Represents the supported package managers.
 */
export enum PackageManager {
  NPM = 'npm',
  Yarn = 'yarn',
  Pip = 'pip',
  Apt = 'apt',
  Brew = 'brew',
}

/**
 * Represents the options for running a script.
 */
export interface RunOptions {
  silent?: boolean;
  verbose?: boolean;
  dryRun?: boolean;
  env?: Record<string, string>;
}

/**
 * Represents a generic key-value pair.
 */
export type KeyValuePair<T = string> = {
  key: string;
  value: T;
};