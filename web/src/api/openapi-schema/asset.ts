/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { Identifier } from "./identifier";

export interface Asset {
  filename: string;
  height: number;
  id: Identifier;
  mime_type: string;
  url: string;
  width: number;
}