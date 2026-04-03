import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { FlatList, FlatItem } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import { useRoute } from '@react-navigation/stack';

const Parser = () => {
  const navigation = useNavigation();
  const route = useRoute();

  const data = [
    { id: 1, title: 'Item 1', description: 'This is item 1' },
    { id: 2, title: 'Item 2', description: 'This is item 2' },
    { id: 3, title: 'Item 3', description: 'This is item 3' },
  ];

  const renderItem = ({ item }) => (
    <FlatItem
      key={item.id}
      title={item.title}
      description={item.description}
      onPress={() => navigation.navigate('Detail', { id: item.id })}
    />
  );

  return (
    <View style={styles.container}>
      <FlatList
        data={data}
        renderItem={renderItem}
        keyExtractor={(item) => item.id.toString()}
        contentContainerStyle={styles.list}
        contentContainerStyle={styles.listContainer}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  list: {
    flex: 1,
    padding: 20,
  },
  listContainer: {
    padding: 20,
  },
});

export default Parser;