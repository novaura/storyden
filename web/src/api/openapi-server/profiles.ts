/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type {
  ProfileGetOKResponse,
  ProfileListOKResponse,
  ProfileListParams,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Query and search profiles.
 */
export type profileListResponse = {
  data: ProfileListOKResponse;
  status: number;
};

export const getProfileListUrl = (params?: ProfileListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/v1/profiles?${normalizedParams.toString()}`;
};

export const profileList = async (
  params?: ProfileListParams,
  options?: RequestInit,
): Promise<profileListResponse> => {
  return fetcher<Promise<profileListResponse>>(getProfileListUrl(params), {
    ...options,
    method: "GET",
  });
};

/**
 * Get a public profile by ID.
 */
export type profileGetResponse = {
  data: ProfileGetOKResponse;
  status: number;
};

export const getProfileGetUrl = (accountHandle: string) => {
  return `/v1/profiles/${accountHandle}`;
};

export const profileGet = async (
  accountHandle: string,
  options?: RequestInit,
): Promise<profileGetResponse> => {
  return fetcher<Promise<profileGetResponse>>(getProfileGetUrl(accountHandle), {
    ...options,
    method: "GET",
  });
};