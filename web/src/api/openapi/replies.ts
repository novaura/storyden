/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";

import type {
  InternalServerErrorResponse,
  NotFoundResponse,
  ReplyCreateBody,
  ReplyCreateOKResponse,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Create a new post within a thread.
 */
export const replyCreate = (
  threadMark: string,
  replyCreateBody: ReplyCreateBody,
) => {
  return fetcher<ReplyCreateOKResponse>({
    url: `/v1/threads/${threadMark}/replies`,
    method: "POST",
    headers: { "Content-Type": "application/json" },
    data: replyCreateBody,
  });
};

export const getReplyCreateMutationFetcher = (threadMark: string) => {
  return (
    _: string,
    { arg }: { arg: ReplyCreateBody },
  ): Promise<ReplyCreateOKResponse> => {
    return replyCreate(threadMark, arg);
  };
};
export const getReplyCreateMutationKey = (threadMark: string) =>
  `/v1/threads/${threadMark}/replies` as const;

export type ReplyCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof replyCreate>>
>;
export type ReplyCreateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useReplyCreate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  threadMark: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof replyCreate>>,
      TError,
      string,
      ReplyCreateBody,
      Awaited<ReturnType<typeof replyCreate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getReplyCreateMutationKey(threadMark);
  const swrFn = getReplyCreateMutationFetcher(threadMark);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
