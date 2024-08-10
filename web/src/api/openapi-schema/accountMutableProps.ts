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
import type { TagListIDs } from "./tagListIDs";

export interface AccountMutableProps {
  bio?: AccountBio;
  handle?: AccountHandle;
  interests?: TagListIDs;
  links?: ProfileExternalLinkList;
  meta?: Metadata;
  name?: AccountName;
}