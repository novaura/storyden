/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  CollectionAddNodeOKResponse,
  CollectionAddPostOKResponse,
  CollectionCreateBody,
  CollectionCreateOKResponse,
  CollectionGetOKResponse,
  CollectionListOKResponse,
  CollectionListParams,
  CollectionRemoveNodeOKResponse,
  CollectionRemovePostOKResponse,
  CollectionUpdateBody,
  CollectionUpdateOKResponse,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Create a collection for curating posts under the authenticated account.

 */
export type collectionCreateResponse = {
  data: CollectionCreateOKResponse;
  status: number;
};

export const getCollectionCreateUrl = () => {
  return `/collections`;
};

export const collectionCreate = async (
  collectionCreateBody: CollectionCreateBody,
  options?: RequestInit,
): Promise<collectionCreateResponse> => {
  return fetcher<Promise<collectionCreateResponse>>(getCollectionCreateUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(collectionCreateBody),
  });
};

/**
 * List all collections using the filtering options.
 */
export type collectionListResponse = {
  data: CollectionListOKResponse;
  status: number;
};

export const getCollectionListUrl = (params?: CollectionListParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value === null) {
      normalizedParams.append(key, "null");
    } else if (value !== undefined) {
      normalizedParams.append(key, value.toString());
    }
  });

  return `/collections?${normalizedParams.toString()}`;
};

export const collectionList = async (
  params?: CollectionListParams,
  options?: RequestInit,
): Promise<collectionListResponse> => {
  return fetcher<Promise<collectionListResponse>>(
    getCollectionListUrl(params),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Get a collection by its ID. Collections can be public or private so the
response will depend on which account is making the request and if the
target collection is public, private, owned or not owned by the account.

 */
export type collectionGetResponse = {
  data: CollectionGetOKResponse;
  status: number;
};

export const getCollectionGetUrl = (collectionMark: string) => {
  return `/collections/${collectionMark}`;
};

export const collectionGet = async (
  collectionMark: string,
  options?: RequestInit,
): Promise<collectionGetResponse> => {
  return fetcher<Promise<collectionGetResponse>>(
    getCollectionGetUrl(collectionMark),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Update a collection owned by the authenticated account.
 */
export type collectionUpdateResponse = {
  data: CollectionUpdateOKResponse;
  status: number;
};

export const getCollectionUpdateUrl = (collectionMark: string) => {
  return `/collections/${collectionMark}`;
};

export const collectionUpdate = async (
  collectionMark: string,
  collectionUpdateBody: CollectionUpdateBody,
  options?: RequestInit,
): Promise<collectionUpdateResponse> => {
  return fetcher<Promise<collectionUpdateResponse>>(
    getCollectionUpdateUrl(collectionMark),
    {
      ...options,
      method: "PATCH",
      body: JSON.stringify(collectionUpdateBody),
    },
  );
};

/**
 * Delete a collection owned by the authenticated account.
 */
export type collectionDeleteResponse = {
  data: void;
  status: number;
};

export const getCollectionDeleteUrl = (collectionMark: string) => {
  return `/collections/${collectionMark}`;
};

export const collectionDelete = async (
  collectionMark: string,
  options?: RequestInit,
): Promise<collectionDeleteResponse> => {
  return fetcher<Promise<collectionDeleteResponse>>(
    getCollectionDeleteUrl(collectionMark),
    {
      ...options,
      method: "DELETE",
    },
  );
};

/**
 * Add a post to a collection. The collection must be owned by the account
making the request. The post can be any published post of any kind.

 */
export type collectionAddPostResponse = {
  data: CollectionAddPostOKResponse;
  status: number;
};

export const getCollectionAddPostUrl = (
  collectionMark: string,
  postId: string,
) => {
  return `/collections/${collectionMark}/posts/${postId}`;
};

export const collectionAddPost = async (
  collectionMark: string,
  postId: string,
  options?: RequestInit,
): Promise<collectionAddPostResponse> => {
  return fetcher<Promise<collectionAddPostResponse>>(
    getCollectionAddPostUrl(collectionMark, postId),
    {
      ...options,
      method: "PUT",
    },
  );
};

/**
 * Remove a post from a collection. The collection must be owned by the
account making the request.

 */
export type collectionRemovePostResponse = {
  data: CollectionRemovePostOKResponse;
  status: number;
};

export const getCollectionRemovePostUrl = (
  collectionMark: string,
  postId: string,
) => {
  return `/collections/${collectionMark}/posts/${postId}`;
};

export const collectionRemovePost = async (
  collectionMark: string,
  postId: string,
  options?: RequestInit,
): Promise<collectionRemovePostResponse> => {
  return fetcher<Promise<collectionRemovePostResponse>>(
    getCollectionRemovePostUrl(collectionMark, postId),
    {
      ...options,
      method: "DELETE",
    },
  );
};

/**
 * Add a node to a collection. The collection must be owned by the account
making the request. The node can be any published node or any node
not published but owned by the collection owner.

 */
export type collectionAddNodeResponse = {
  data: CollectionAddNodeOKResponse;
  status: number;
};

export const getCollectionAddNodeUrl = (
  collectionMark: string,
  nodeId: string,
) => {
  return `/collections/${collectionMark}/nodes/${nodeId}`;
};

export const collectionAddNode = async (
  collectionMark: string,
  nodeId: string,
  options?: RequestInit,
): Promise<collectionAddNodeResponse> => {
  return fetcher<Promise<collectionAddNodeResponse>>(
    getCollectionAddNodeUrl(collectionMark, nodeId),
    {
      ...options,
      method: "PUT",
    },
  );
};

/**
 * Remove a node from a collection. The collection must be owned by the
account making the request.

 */
export type collectionRemoveNodeResponse = {
  data: CollectionRemoveNodeOKResponse;
  status: number;
};

export const getCollectionRemoveNodeUrl = (
  collectionMark: string,
  nodeId: string,
) => {
  return `/collections/${collectionMark}/nodes/${nodeId}`;
};

export const collectionRemoveNode = async (
  collectionMark: string,
  nodeId: string,
  options?: RequestInit,
): Promise<collectionRemoveNodeResponse> => {
  return fetcher<Promise<collectionRemoveNodeResponse>>(
    getCollectionRemoveNodeUrl(collectionMark, nodeId),
    {
      ...options,
      method: "DELETE",
    },
  );
};
