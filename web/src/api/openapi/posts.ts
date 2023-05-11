/**
 * Generated by orval v6.14.3 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type {
  PostCreateOKResponse,
  PostCreateBody,
  PostUpdateOKResponse,
  PostUpdateBody,
  PostReactAddOKResponse,
  PostReactAddBody,
} from "./schemas";
import { fetcher } from "../client";

/**
 * Create a new post within a thread.
 */
export const postCreate = (
  threadMark: string,
  postCreateBody: PostCreateBody
) => {
  return fetcher<PostCreateOKResponse>({
    url: `/v1/threads/${threadMark}/posts`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: postCreateBody,
  });
};

/**
 * Publish changes to a single post.
 */
export const postUpdate = (postId: string, postUpdateBody: PostUpdateBody) => {
  return fetcher<PostUpdateOKResponse>({
    url: `/v1/posts/${postId}`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: postUpdateBody,
  });
};

/**
 * Add a reaction to a post.
 */
export const postReactAdd = (
  postId: string,
  postReactAddBody: PostReactAddBody
) => {
  return fetcher<PostReactAddOKResponse>({
    url: `/v1/posts/${postId}/reacts`,
    method: "put",
    headers: { "Content-Type": "application/json" },
    data: postReactAddBody,
  });
};
