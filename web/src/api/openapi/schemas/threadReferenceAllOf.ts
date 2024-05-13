/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AssetList } from "./assetList";
import type { CategoryReference } from "./categoryReference";
import type { CollectionList } from "./collectionList";
import type { Link } from "./link";
import type { Metadata } from "./metadata";
import type { ProfileReference } from "./profileReference";
import type { ReactList } from "./reactList";
import type { ThreadMark } from "./threadMark";

export type ThreadReferenceAllOf = {
  assets: AssetList;
  author: ProfileReference;
  category: CategoryReference;
  collections: CollectionList;
  link?: Link;
  meta?: Metadata;
  /** Whether the thread is pinned in this category. */
  pinned: boolean;
  /** The number of posts under this thread. */
  readonly post_count: number;
  reacts: ReactList;
  /** A short version of the thread's body text for use in previews.
   */
  readonly short: string;
  slug: ThreadMark;
  /** A list of tags associated with the thread. */
  tags: string[];
  /** The title of the thread. */
  title: string;
};
