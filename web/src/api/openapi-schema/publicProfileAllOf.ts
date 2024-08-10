/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AccountBio } from "./accountBio";
import type { AccountHandle } from "./accountHandle";
import type { AccountName } from "./accountName";
import type { Metadata } from "./metadata";
import type { ProfileExternalLinkList } from "./profileExternalLinkList";
import type { TagList } from "./tagList";

export type PublicProfileAllOf = {
  bio: AccountBio;
  createdAt: string;
  handle: AccountHandle;
  image?: string;
  interests: TagList;
  links: ProfileExternalLinkList;
  meta: Metadata;
  name: AccountName;
};