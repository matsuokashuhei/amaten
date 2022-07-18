import { credentials } from "grpc";
import { CustomerServiceClient } from "../../proto/grocery_grpc_pb";
import {
  CreateCustomerRequest,
  DeleteCustomerRequest,
  GetCustomerRequest,
} from "../../proto/grocery_pb";
import { MutationResolvers, QueryResolvers } from "../types";

export const customer: QueryResolvers["customer"] = async (
  parent,
  args,
  context,
  info
) => {
  const client = new CustomerServiceClient(
    "customer-service:50051",
    credentials.createInsecure()
  );
  const req = new GetCustomerRequest();
  req.setUid(args.uid);
  return new Promise((resolve, reject) => {
    client.getCustomer(req, (error, res) => {
      if (error) {
        reject(error);
      } else {
        resolve({ uid: res.getUid(), name: res.getName() });
      }
    });
  });
};

export const createCustomer: MutationResolvers["createCustomer"] = async (
  parent,
  args,
  context,
  info
) => {
  const client = new CustomerServiceClient(
    "customer-service:50051",
    credentials.createInsecure()
  );
  const req = new CreateCustomerRequest();
  req.setName(args.input.name);
  req.setEmail(args.input.email);
  req.setPassword(args.input.password);
  return new Promise((resolve, reject) => {
    client.createCustomer(req, (error, res) => {
      if (error) {
        reject(error);
      } else {
        resolve({ uid: res.getUid(), name: res.getName() });
      }
    });
  });
};

export const deleteCustomer: MutationResolvers["deleteCustomer"] = async (
  parent,
  args,
  context,
  info
) => {
  const client = new CustomerServiceClient(
    "customer-service:50051",
    credentials.createInsecure()
  );
  const req = new DeleteCustomerRequest();
  req.setUid(args.input.uid);
  return new Promise((resolve, reject) => {
    client.deleteCustomer(req, (error, res) => {
      if (error) {
        reject(error);
      } else {
        resolve({ uid: res.getUid(), name: res.getName() });
      }
    });
  });
};
