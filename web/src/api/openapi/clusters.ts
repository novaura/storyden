/**
 * Generated by orval v6.24.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Arguments, Key, SWRConfiguration } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";

import type {
  ClusterAddChildOKResponse,
  ClusterAddItemOKResponse,
  ClusterCreateBody,
  ClusterCreateOKResponse,
  ClusterDeleteOKResponse,
  ClusterDeleteParams,
  ClusterGetOKResponse,
  ClusterListOKResponse,
  ClusterListParams,
  ClusterRemoveChildOKResponse,
  ClusterRemoveItemOKResponse,
  ClusterUpdateBody,
  ClusterUpdateOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
  VisibilityUpdateBody,
} from "./schemas";

/**
 * Create a cluster for grouping items and other clusters together.

 */
export const clusterCreate = (clusterCreateBody: ClusterCreateBody) => {
  return fetcher<ClusterCreateOKResponse>({
    url: `/v1/clusters`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: clusterCreateBody,
  });
};

export const getClusterCreateMutationFetcher = () => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterCreateOKResponse> => {
    return clusterCreate(arg as ClusterCreateBody);
  };
};
export const getClusterCreateMutationKey = () => `/v1/clusters` as const;

export type ClusterCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterCreate>>
>;
export type ClusterCreateMutationError =
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useClusterCreate = <
  TError = UnauthorisedResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof clusterCreate>>,
    TError,
    string,
    Arguments,
    Awaited<ReturnType<typeof clusterCreate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getClusterCreateMutationKey();
  const swrFn = getClusterCreateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * List clusters using the given filters. Can be used to get a full tree.

 */
export const clusterList = (params?: ClusterListParams) => {
  return fetcher<ClusterListOKResponse>({
    url: `/v1/clusters`,
    method: "GET",
    params,
  });
};

export const getClusterListKey = (params?: ClusterListParams) =>
  [`/v1/clusters`, ...(params ? [params] : [])] as const;

export type ClusterListQueryResult = NonNullable<
  Awaited<ReturnType<typeof clusterList>>
>;
export type ClusterListQueryError =
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(
  params?: ClusterListParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof clusterList>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getClusterListKey(params) : null));
  const swrFn = () => clusterList(params);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Get a cluster by its URL slug.
 */
export const clusterGet = (clusterSlug: string) => {
  return fetcher<ClusterGetOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "GET",
  });
};

export const getClusterGetKey = (clusterSlug: string) =>
  [`/v1/clusters/${clusterSlug}`] as const;

export type ClusterGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof clusterGet>>
>;
export type ClusterGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof clusterGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!clusterSlug;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getClusterGetKey(clusterSlug) : null));
  const swrFn = () => clusterGet(clusterSlug);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
/**
 * Update a cluster.
 */
export const clusterUpdate = (
  clusterSlug: string,
  clusterUpdateBody: ClusterUpdateBody,
) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: clusterUpdateBody,
  });
};

export const getClusterUpdateMutationFetcher = (clusterSlug: string) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterUpdateOKResponse> => {
    return clusterUpdate(clusterSlug, arg as ClusterUpdateBody);
  };
};
export const getClusterUpdateMutationKey = (clusterSlug: string) =>
  `/v1/clusters/${clusterSlug}` as const;

export type ClusterUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterUpdate>>
>;
export type ClusterUpdateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterUpdate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterUpdate>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterUpdate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getClusterUpdateMutationKey(clusterSlug);
  const swrFn = getClusterUpdateMutationFetcher(clusterSlug);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Delete a cluster and move all children to its parent or root.
 */
export const clusterDelete = (
  clusterSlug: string,
  params?: ClusterDeleteParams,
) => {
  return fetcher<ClusterDeleteOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "DELETE",
    params,
  });
};

export const getClusterDeleteMutationFetcher = (
  clusterSlug: string,
  params?: ClusterDeleteParams,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterDeleteOKResponse> => {
    return clusterDelete(clusterSlug, params);
  };
};
export const getClusterDeleteMutationKey = (clusterSlug: string) =>
  `/v1/clusters/${clusterSlug}` as const;

export type ClusterDeleteMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterDelete>>
>;
export type ClusterDeleteMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterDelete = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  params?: ClusterDeleteParams,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterDelete>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterDelete>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getClusterDeleteMutationKey(clusterSlug);
  const swrFn = getClusterDeleteMutationFetcher(clusterSlug, params);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Update the visibility cluster.
 */
export const clusterUpdateVisibility = (
  clusterSlug: string,
  visibilityUpdateBody: VisibilityUpdateBody,
) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}/visibility`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: visibilityUpdateBody,
  });
};

export const getClusterUpdateVisibilityMutationFetcher = (
  clusterSlug: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterUpdateOKResponse> => {
    return clusterUpdateVisibility(clusterSlug, arg as VisibilityUpdateBody);
  };
};
export const getClusterUpdateVisibilityMutationKey = (clusterSlug: string) =>
  `/v1/clusters/${clusterSlug}/visibility` as const;

export type ClusterUpdateVisibilityMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterUpdateVisibility>>
>;
export type ClusterUpdateVisibilityMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterUpdateVisibility = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterUpdateVisibility>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterUpdateVisibility>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getClusterUpdateVisibilityMutationKey(clusterSlug);
  const swrFn = getClusterUpdateVisibilityMutationFetcher(clusterSlug);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add an asset to a cluster.
 */
export const clusterAddAsset = (clusterSlug: string, assetId: string) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}/assets/${assetId}`,
    method: "PUT",
  });
};

export const getClusterAddAssetMutationFetcher = (
  clusterSlug: string,
  assetId: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterUpdateOKResponse> => {
    return clusterAddAsset(clusterSlug, assetId);
  };
};
export const getClusterAddAssetMutationKey = (
  clusterSlug: string,
  assetId: string,
) => `/v1/clusters/${clusterSlug}/assets/${assetId}` as const;

export type ClusterAddAssetMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterAddAsset>>
>;
export type ClusterAddAssetMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterAddAsset = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  assetId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterAddAsset>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterAddAsset>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getClusterAddAssetMutationKey(clusterSlug, assetId);
  const swrFn = getClusterAddAssetMutationFetcher(clusterSlug, assetId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove an asset from a cluster.
 */
export const clusterRemoveAsset = (clusterSlug: string, assetId: string) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}/assets/${assetId}`,
    method: "DELETE",
  });
};

export const getClusterRemoveAssetMutationFetcher = (
  clusterSlug: string,
  assetId: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterUpdateOKResponse> => {
    return clusterRemoveAsset(clusterSlug, assetId);
  };
};
export const getClusterRemoveAssetMutationKey = (
  clusterSlug: string,
  assetId: string,
) => `/v1/clusters/${clusterSlug}/assets/${assetId}` as const;

export type ClusterRemoveAssetMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterRemoveAsset>>
>;
export type ClusterRemoveAssetMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterRemoveAsset = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  assetId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterRemoveAsset>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterRemoveAsset>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getClusterRemoveAssetMutationKey(clusterSlug, assetId);
  const swrFn = getClusterRemoveAssetMutationFetcher(clusterSlug, assetId);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Set a cluster's parent to the specified cluster
 */
export const clusterAddCluster = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return fetcher<ClusterAddChildOKResponse>({
    url: `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}`,
    method: "PUT",
  });
};

export const getClusterAddClusterMutationFetcher = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterAddChildOKResponse> => {
    return clusterAddCluster(clusterSlug, clusterSlugChild);
  };
};
export const getClusterAddClusterMutationKey = (
  clusterSlug: string,
  clusterSlugChild: string,
) => `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}` as const;

export type ClusterAddClusterMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterAddCluster>>
>;
export type ClusterAddClusterMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterAddCluster = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  clusterSlugChild: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterAddCluster>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterAddCluster>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getClusterAddClusterMutationKey(clusterSlug, clusterSlugChild);
  const swrFn = getClusterAddClusterMutationFetcher(
    clusterSlug,
    clusterSlugChild,
  );

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove a cluster from its parent cluster and back to the top level.

 */
export const clusterRemoveCluster = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return fetcher<ClusterRemoveChildOKResponse>({
    url: `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}`,
    method: "DELETE",
  });
};

export const getClusterRemoveClusterMutationFetcher = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterRemoveChildOKResponse> => {
    return clusterRemoveCluster(clusterSlug, clusterSlugChild);
  };
};
export const getClusterRemoveClusterMutationKey = (
  clusterSlug: string,
  clusterSlugChild: string,
) => `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}` as const;

export type ClusterRemoveClusterMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterRemoveCluster>>
>;
export type ClusterRemoveClusterMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterRemoveCluster = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  clusterSlugChild: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterRemoveCluster>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterRemoveCluster>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getClusterRemoveClusterMutationKey(clusterSlug, clusterSlugChild);
  const swrFn = getClusterRemoveClusterMutationFetcher(
    clusterSlug,
    clusterSlugChild,
  );

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Add an item to a cluster.
 */
export const clusterAddItem = (clusterSlug: string, itemSlug: string) => {
  return fetcher<ClusterAddItemOKResponse>({
    url: `/v1/clusters/${clusterSlug}/items/${itemSlug}`,
    method: "PUT",
  });
};

export const getClusterAddItemMutationFetcher = (
  clusterSlug: string,
  itemSlug: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterAddItemOKResponse> => {
    return clusterAddItem(clusterSlug, itemSlug);
  };
};
export const getClusterAddItemMutationKey = (
  clusterSlug: string,
  itemSlug: string,
) => `/v1/clusters/${clusterSlug}/items/${itemSlug}` as const;

export type ClusterAddItemMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterAddItem>>
>;
export type ClusterAddItemMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterAddItem = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  itemSlug: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterAddItem>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterAddItem>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getClusterAddItemMutationKey(clusterSlug, itemSlug);
  const swrFn = getClusterAddItemMutationFetcher(clusterSlug, itemSlug);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Remove an item from a cluster.
 */
export const clusterRemoveItem = (clusterSlug: string, itemSlug: string) => {
  return fetcher<ClusterRemoveItemOKResponse>({
    url: `/v1/clusters/${clusterSlug}/items/${itemSlug}`,
    method: "DELETE",
  });
};

export const getClusterRemoveItemMutationFetcher = (
  clusterSlug: string,
  itemSlug: string,
) => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<ClusterRemoveItemOKResponse> => {
    return clusterRemoveItem(clusterSlug, itemSlug);
  };
};
export const getClusterRemoveItemMutationKey = (
  clusterSlug: string,
  itemSlug: string,
) => `/v1/clusters/${clusterSlug}/items/${itemSlug}` as const;

export type ClusterRemoveItemMutationResult = NonNullable<
  Awaited<ReturnType<typeof clusterRemoveItem>>
>;
export type ClusterRemoveItemMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterRemoveItem = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  itemSlug: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof clusterRemoveItem>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof clusterRemoveItem>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ??
    getClusterRemoveItemMutationKey(clusterSlug, itemSlug);
  const swrFn = getClusterRemoveItemMutationFetcher(clusterSlug, itemSlug);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
