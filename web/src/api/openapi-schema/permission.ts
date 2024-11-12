/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */

export type Permission = (typeof Permission)[keyof typeof Permission];

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const Permission = {
  CREATE_POST: "CREATE_POST",
  READ_PUBLISHED_THREADS: "READ_PUBLISHED_THREADS",
  CREATE_REACTION: "CREATE_REACTION",
  MANAGE_POSTS: "MANAGE_POSTS",
  MANAGE_CATEGORIES: "MANAGE_CATEGORIES",
  CREATE_INVITATION: "CREATE_INVITATION",
  READ_PUBLISHED_LIBRARY: "READ_PUBLISHED_LIBRARY",
  MANAGE_LIBRARY: "MANAGE_LIBRARY",
  SUBMIT_LIBRARY_NODE: "SUBMIT_LIBRARY_NODE",
  UPLOAD_ASSET: "UPLOAD_ASSET",
  MANAGE_EVENTS: "MANAGE_EVENTS",
  LIST_PROFILES: "LIST_PROFILES",
  READ_PROFILE: "READ_PROFILE",
  CREATE_COLLECTION: "CREATE_COLLECTION",
  LIST_COLLECTIONS: "LIST_COLLECTIONS",
  READ_COLLECTION: "READ_COLLECTION",
  MANAGE_COLLECTIONS: "MANAGE_COLLECTIONS",
  COLLECTION_SUBMIT: "COLLECTION_SUBMIT",
  MANAGE_SETTINGS: "MANAGE_SETTINGS",
  MANAGE_SUSPENSIONS: "MANAGE_SUSPENSIONS",
  MANAGE_ROLES: "MANAGE_ROLES",
  ADMINISTRATOR: "ADMINISTRATOR",
} as const;
