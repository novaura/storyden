/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  LikePostGetOKResponse,
  LikeProfileGetOKResponse,
  LikeProfileGetParams,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Retreives all likes for the given post. Not paginated (yet.)
 */
export type likePostGetResponse = {
  data: LikePostGetOKResponse;
  status: number;
};

export const getLikePostGetUrl = (postId: string) => {
  return `/likes/posts/${postId}`;
};

export const likePostGet = async (
  postId: string,
  options?: RequestInit,
): Promise<likePostGetResponse> => {
  return fetcher<Promise<likePostGetResponse>>(getLikePostGetUrl(postId), {
    ...options,
    method: "GET",
  });
};

/**
 * Add a like/vote to a post. A "like" is pretty much what you'd expect for
any modern social platform, it will inform the feed algorithm and the
account's recommendations as well as listing the post on their profile.
Idempotent operation where repeated use will do nothing.

 */
export type likePostAddResponse = {
  data: void;
  status: number;
};

export const getLikePostAddUrl = (postId: string) => {
  return `/likes/posts/${postId}`;
};

export const likePostAdd = async (
  postId: string,
  options?: RequestInit,
): Promise<likePostAddResponse> => {
  return fetcher<Promise<likePostAddResponse>>(getLikePostAddUrl(postId), {
    ...options,
    method: "PUT",
  });
};

/**
 * Removes a like/vote from the authenticated account for the post. It will
perform the inverse of any changes to the account's algorithm. Also is
idempotent, so repeated use will do nothing after being actioned once.

 */
export type likePostRemoveResponse = {
  data: void;
  status: number;
};

export const getLikePostRemoveUrl = (postId: string) => {
  return `/likes/posts/${postId}`;
};

export const likePostRemove = async (
  postId: string,
  options?: RequestInit,
): Promise<likePostRemoveResponse> => {
  return fetcher<Promise<likePostRemoveResponse>>(
    getLikePostRemoveUrl(postId),
    {
      ...options,
      method: "DELETE",
    },
  );
};

/**
 * Retreives all the likes that the given profile has given.
 */
export type likeProfileGetResponse = {
  data: LikeProfileGetOKResponse;
  status: number;
};

export const getLikeProfileGetUrl = (
  accountHandle: string,
  params?: LikeProfileGetParams,
) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/likes/profiles/${accountHandle}?${normalizedParams.toString()}`;
};

export const likeProfileGet = async (
  accountHandle: string,
  params?: LikeProfileGetParams,
  options?: RequestInit,
): Promise<likeProfileGetResponse> => {
  return fetcher<Promise<likeProfileGetResponse>>(
    getLikeProfileGetUrl(accountHandle, params),
    {
      ...options,
      method: "GET",
    },
  );
};