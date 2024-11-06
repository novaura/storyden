/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  TagGetOKResponse,
  TagListOKResponse,
  TagListParams,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Get a list of all tags on the site.
 */
export type tagListResponse = {
  data: TagListOKResponse;
  status: number;
};

export const getTagListUrl = (params?: TagListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/tags?${normalizedParams.toString()}`;
};

export const tagList = async (
  params?: TagListParams,
  options?: RequestInit,
): Promise<tagListResponse> => {
  return fetcher<Promise<tagListResponse>>(getTagListUrl(params), {
    ...options,
    method: "GET",
  });
};

/**
 * Get information about a tag.
 */
export type tagGetResponse = {
  data: TagGetOKResponse;
  status: number;
};

export const getTagGetUrl = (tagName: string) => {
  return `/tags/${tagName}`;
};

export const tagGet = async (
  tagName: string,
  options?: RequestInit,
): Promise<tagGetResponse> => {
  return fetcher<Promise<tagGetResponse>>(getTagGetUrl(tagName), {
    ...options,
    method: "GET",
  });
};