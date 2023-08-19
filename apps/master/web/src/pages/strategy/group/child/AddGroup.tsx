import React, { useEffect } from "react";
import EditGroup from "@/pages/strategy/group/child/EditGroup";
import { GroupCreateItem } from "@/apis/prom/group/group";

export interface AddGroupProps {
  children?: React.ReactNode;
  title?: React.ReactNode;
  onFinished?: () => void;
  initialValues?: GroupCreateItem;
  groupId?: number;
  setRefreshLock?: React.Dispatch<React.SetStateAction<boolean>>;
}

const defaultAddGroupName = "新建分组";

const AddGroup: React.FC<AddGroupProps> = (props) => {
  const {
    children = defaultAddGroupName,
    title = defaultAddGroupName,
    onFinished,
    initialValues,
    groupId,
    setRefreshLock,
  } = props;
  const [visible, setVisible] = React.useState<boolean>(false);

  const handleOpenEditGroup = () => {
    setVisible(true);
  };

  const handleOnClose = () => {
    setVisible(false);
  };

  const handleOnFinished = () => {
    setVisible(false);
    onFinished?.();
  };

  useEffect(() => {
    setRefreshLock?.(visible);
  }, []);

  return (
    <>
      <div onClick={handleOpenEditGroup}>{children}</div>
      <EditGroup
        title={title}
        visible={visible}
        onClose={handleOnClose}
        onFinished={handleOnFinished}
        initialValues={initialValues}
        id={groupId}
      />
    </>
  );
};

export default AddGroup;
